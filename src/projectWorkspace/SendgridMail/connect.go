package SendgridMail

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
)

type MailingClient struct {
	ISGClient SendInterface
}

func (m *MailingClient) Connect(appKey string) bool {

	client := sendgrid.NewSendClient(appKey)

	m.ISGClient = AdaptSClient(client)

	if m.ISGClient != nil {
		log.Println("Sendgrid Client initialized successfully")
		return true
	}
	log.Printf("SendGrid Client creation failed")
	return false
}


