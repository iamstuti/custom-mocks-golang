package InterfacesAndMocks

import (
	"projectWorkspace/projectWorkspace/ExternalServices/PaymentServices/models"
)

type PaymentInterface interface {
	GetUserPaymentStatus(string,string) (models.UserPayment, error)
	CheckPaymentStatus(string,string)(string,error)
	GetUserPaymentHistory(string)([]models.UserPayment,error)
}
