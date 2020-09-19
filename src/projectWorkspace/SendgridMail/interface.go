package SendgridMail

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sendgrid/rest"
)

//Client is the interface that wraps a sendgrid.Client.
type SendInterface interface {
	Send(email *mail.SGMailV3) (*rest.Response, error)
}