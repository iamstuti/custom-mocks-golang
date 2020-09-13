package Sendgrid-Mail

//Client is the interface that wraps a sendgrid.Client.
type SGClient interface {
	Send(email *mail.SGMailV3) (*rest.Response, error)
}