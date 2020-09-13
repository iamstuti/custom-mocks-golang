package Sendgrid-Mail

type sgClient struct {
	*sendgrid.Client 
}

func AdaptSGClient(sg *sendgrid.Client) SGClient{
	sgClient{sg}
}

func (sc sgClient) Send(email *mail.SGMailV3) (*rest.Response, error){
	sc.Client.Send(email)
}
