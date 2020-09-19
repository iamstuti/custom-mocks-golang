package SendgridMail
import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/sendgrid/sendgrid-go"
)

type sgClient struct {
	*sendgrid.Client 
}

func AdaptSClient(sg *sendgrid.Client) SendInterface{
	return sgClient{sg}
}

func (sc sgClient) Send(email *mail.SGMailV3) (*rest.Response, error){
	return sc.Client.Send(email)
}
