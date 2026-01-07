package application

import (
	"github.com/akkaraponph/email-job-temporal/internal/adapters/http/handlers"
	"github.com/akkaraponph/email-job-temporal/internal/adapters/http/routers"
	"github.com/akkaraponph/email-job-temporal/internal/core/services"
	"go.temporal.io/sdk/client"
	"gofr.dev/pkg/gofr"
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
