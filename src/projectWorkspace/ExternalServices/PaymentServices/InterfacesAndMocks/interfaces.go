package InterfacesAndMocks

import (
	"projectWorkspace/projectWorkspace/ExternalServices/PaymentServices/models"
)

type PaymentStatusInterface interface {
	GetUserPaymentStatus(string,string) (models.UserPayment, error)
}

type CheckPaymentInterface interface {
	CheckPaymentStatus(string,string)(string,models.UserPayment,error)
}

type PaymentHistoryInterface interface {
	GetUserPaymentHistory(string)([]models.UserPayment,error)
}






