package configs

import (
	"fmt"
	"log"
	"os"

	"gofr.dev/pkg/gofr"
)

type GoFrHttpServiceParams struct {
	Port    string
	Address string
}

func NewGoFrHttpServiceParams() *GoFrHttpServiceParams {
	fmt.Println("GOFR HTTP PORT = ", SERVER_HTTP_PORT)
	return &GoFrHttpServiceParams{
		Port:    SERVER_HTTP_PORT,
		Address: "",
	}
}

func NewGoFrHTTPService(params *GoFrHttpServiceParams) *gofr.App {
	// Set port in environment for GoFr to pick up
	if params.Port != "" {
		os.Setenv("HTTP_PORT", params.Port)
	}

	// GoFr automatically handles observability (logs, metrics, traces)
	// and REST standards, so we don't need to configure logger middleware
	app := gofr.New()

	// Note: CORS can be configured via environment variables or custom middleware if needed
	// GoFr follows REST standards by default

	log.Printf("Launching GoFr HTTP listener on port [%s]...", params.Port)

	return app
}
