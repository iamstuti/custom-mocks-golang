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





type PaymentInterface interface {
	GetUserPaymentStatus(string,string) (models.UserPayment, error)
	CheckPaymentStatus(string,string)(string,error)
	GetUserPaymentHistory(string)([]models.UserPayment,error)
}





