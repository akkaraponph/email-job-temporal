package handlers

import (
	"fmt"

	"github.com/akkaraponph/email-job-temporal/internal/core/domain"
	"github.com/akkaraponph/email-job-temporal/internal/core/ports"
	"github.com/akkaraponph/email-job-temporal/pkg/configs"
	"gofr.dev/pkg/gofr"
)

type (
	IEmailHandler interface {
		HandleSendEmail(c *gofr.Context) (interface{}, error)
	}
	EmailHandlerImpls struct {
		emailSrv ports.IEmailService
	}
)

func NewEmailHandler(
	emailSrv ports.IEmailService,
) IEmailHandler {
	return &EmailHandlerImpls{emailSrv: emailSrv}
}

// SendEmail implements IEmailHandler.
func (e *EmailHandlerImpls) HandleSendEmail(c *gofr.Context) (interface{}, error) {
	var emailRequest domain.EmailDto
	// Parse the request body into the struct
	if err := c.Bind(&emailRequest); err != nil {
		return nil, fmt.Errorf("unable to parse request body: %w", err)
	}
	if emailRequest.Sender == "" {
		emailRequest.Sender = configs.SMTP_SENDER
	}
	err := e.emailSrv.SendEmail(domain.EmailDto{
		Sender:       emailRequest.Sender,
		Receiver:     emailRequest.Receiver,
		Subject:      emailRequest.Subject,
		HTMLTemplate: emailRequest.HTMLTemplate,
		CC:           emailRequest.CC,
	})
	if err != nil {
		return nil, err
	}
	return map[string]string{"status": "success"}, nil
}
