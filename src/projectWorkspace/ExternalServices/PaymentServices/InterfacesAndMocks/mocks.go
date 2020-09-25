package InterfacesAndMocks

import (
	"errors"
	"time"
	"projectWorkspace/projectWorkspace/ExternalServices/PaymentServices/models"
)

type MockCheckPayment struct {
	CheckPaymentInterface
}

type MockPaymentStatus struct {
	PaymentStatusInterface
}

const Err_Payment = "User Payment Status unavailable"
const Failed_Payment = "Payment failed"

func (mcp MockCheckPayment)	CheckPaymentStatus(userId string,accountId string)(string,models.UserPayment,error){
	paymentObj := models.UserPayment{
		ReferenceNumber: 952634179,
		PaymentDate:     time.Now(),
		PaymentMode:     "Citibank Credit Card",
		Amount:          50.25,
		Status:          "success",
	}


	return "Payment successful",paymentObj,nil
}


func (mp MockPaymentStatus) GetUserPaymentStatus(userId string,accountId string) (models.UserPayment, error) {
	userObj := models.UserPayment{
		ReferenceNumber: 121020797,
		PaymentDate:     time.Now(),
		PaymentMode:     "American Express Credit Card",
		Amount:          50.25,
		Status:          "success",
	}

	return userObj, nil
}

type MockPaymentStatusErr struct{
	PaymentStatusInterface
}

func (mpr MockPaymentStatusErr) GetUserPaymentStatus(userId string,accountId string) (models.UserPayment, error) {
	
	return models.UserPayment{}, errors.New(Err_Payment)
}

type MockCheckPaymentError struct {
	CheckPaymentInterface
}

func (mcp MockCheckPaymentError)CheckPaymentStatus(userId string,accountId string)(string,models.UserPayment,error){
	return "",models.UserPayment{},errors.New(Failed_Payment)
}






