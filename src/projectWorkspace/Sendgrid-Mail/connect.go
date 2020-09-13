package Sendgrid-Mail

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
)

type MailingClient struct {
	ISGClient SGlient
}

func (m *MailingClient) Connect(appKey string) bool {

	client := sendgrid.NewSendClient(appKey)

	m.ISGClient = AdaptSClient(client)

	if m.SGClient != nil {
		log.Println("Sendgrid Client initialized successfully")
		return true
	}
	log.Printf("SendGrid Client creation failed")
	return false
}


