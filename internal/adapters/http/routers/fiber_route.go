package routers

import "github.com/billowdev/email-job-temporal/internal/adapters/http/handlers"

func (r RouterImpls) CreateEmailRoute(h handlers.IEmailHandler) {
	r.app.POST("/v1/emails/send", h.HandleSendEmail)
}
