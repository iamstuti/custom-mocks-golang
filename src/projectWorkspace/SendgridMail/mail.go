package SendgridMail

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"projectWorkspace/projectWorkspace/model"
	"log"
)

func TriggerMail(mailParams interface{},client SendInterface) error{
	userParams := mailParams.(model.User)

	contentStr := "This is to confirm that " + userParams.UserName + "has been successfully added to the account : "+ userParams.AccountID

	mailContent := mail.NewContent("text/html",contentStr)

	from := mail.NewEmail("test123","test123@abc.com")

	to := mail.NewEmail("receiver123","all-receiver@abc.com")


	m := mail.NewV3MailInit(from,"Test Email",to,mailContent) 
	
	respMail,errRespMail := client.Send(m)

	log.Println(respMail)

	return errRespMail
}