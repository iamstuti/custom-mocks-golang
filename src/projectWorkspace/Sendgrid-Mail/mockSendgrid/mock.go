package mockSendgrid

import (
	"errors"
	SG "projectWorkspace/Sendgrid-Mail"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	ErrSendMail string = "Error in sending mail"
)


type MockSendgridClient struct {
	SG.SGClient
}

func (mockObj MockSendgridClient)Send(email *mail.SGMailV3) (*rest.Response, error) {
	return nil, nil
}

type MockSendgridClientErr struct{
	SG.SGClient	
}

func (mockObjErr MockSendgridClientErr)Send(email *mail.SGMailV3) (*rest.Response, error) {
	return nil, errors.New(ErrSendMail)
}



