package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/akkaraponph/email-job-temporal/cmd/application"
	"github.com/akkaraponph/email-job-temporal/internal/adapters/temporal/worker"
	"github.com/akkaraponph/email-job-temporal/pkg/configs"
	"go.temporal.io/sdk/client"
	temporalLog "go.temporal.io/sdk/log"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := temporalLog.NewStructuredLogger(slog.Default())
	hostPort := client.DefaultHostPort
	if configs.TEMPORAL_CLIENT_URL != "" {
		hostPort = configs.TEMPORAL_CLIENT_URL
	}

	temporalClient, err := client.Dial(client.Options{
		HostPort:  hostPort,
		Namespace: configs.TEMPORAL_NAMESPACE,
		Logger:    logger,
	})
	if err != nil {
		log.Fatal("Failed to create Temporal client:", err)
	}
	defer temporalClient.Close()

	var wg sync.WaitGroup

	// Start Temporal worker in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker.RegisterTemporalWorkflow(temporalClient)
	}()

	// Initialize HTTP application
	params := configs.NewGoFrHttpServiceParams()
	gofrApp := configs.NewGoFrHTTPService(params)
	httpApp := application.AppContainer(gofrApp, temporalClient)

	// Start HTTP application in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpApp.Run()
	}()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Wait for shutdown signal
	sig := <-sigChan
	log.Printf("Received signal: %v, shutting down...", sig)

	// Cancel context to signal shutdown
	cancel()

	// Close Temporal client (this will stop the worker)
	temporalClient.Close()

	// Wait for goroutines to finish
	wg.Wait()

	log.Println("Shutdown complete")
}
