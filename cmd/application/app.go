package application

import (
	"github.com/billowdev/email-job-temporal/internal/adapters/http/handlers"
	"github.com/billowdev/email-job-temporal/internal/adapters/http/routers"
	"github.com/billowdev/email-job-temporal/internal/core/services"
	"gofr.dev/pkg/gofr"
	"go.temporal.io/sdk/client"
)

func AppContainer(app *gofr.App, temporalClient client.Client) *gofr.App {
	route := routers.NewRoute(app)
	EmailApp(route, temporalClient)
	return app
}

func EmailApp(r routers.RouterImpls, temporalClient client.Client) {
	emailSrv := services.NewEmailService(temporalClient)
	emailhandlers := handlers.NewEmailHandler(emailSrv)
	r.CreateEmailRoute(emailhandlers)
}
