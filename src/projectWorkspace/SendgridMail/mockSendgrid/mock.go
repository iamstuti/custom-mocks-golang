package mockSendgrid

import (
	"errors"
	SG "projectWorkspace/projectWorkspace/SendgridMail"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	ErrSendMail string = "Error in sending mail"
)


type MockSendgridClient struct {
	SG.SendInterface
}

func (mockObj MockSendgridClient)Send(email *mail.SGMailV3) (*rest.Response, error) {
	return nil, nil
}

type MockSendgridClientErr struct{
	SG.SendInterface	
}

func (mockObjErr MockSendgridClientErr)Send(email *mail.SGMailV3) (*rest.Response, error) {
	return nil, errors.New(ErrSendMail)
}



