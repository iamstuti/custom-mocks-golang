package PaymentServices

import (
	"encoding/json"
	"net/url"
	"projectWorkspace/projectWorkspace/ExternalServices/PaymentServices/models"
	"errors"
	"projectWorkspace/projectWorkspace/ExternalServices/PaymentServices/PaymentClientUtils"
	"projectWorkspace/projectWorkspace/ExternalServices/PaymentServices/InterfacesAndMocks"
	"strings"
)

const PAYMENT_SUCCESS = "success"
const PAYMENT_FAILURE = "failure"
const paymentUrl = "locahost:9000/api/v1/paymentStatus?"

type CheckPaymentStatus struct {
	InterfacesAndMocks.CheckPaymentInterface
	PaymentStatusParams
}

type  PaymentStatusParams struct {
	InterfacesAndMocks.PaymentStatusInterface
	PaymentClientUtils.ClientInterface
}

func (cp CheckPaymentStatus)CheckPaymentStatus(userId string,accountId string) (string, error){

	paymentObj, errObj := cp.PaymentStatusParams.GetUserPaymentStatus(userId,accountId)

	if errObj !=nil {
		return "", errors.New("Error occurred in fetching user's payment status")
	}
	
	if !strings.EqualFold(paymentObj.Status, PAYMENT_SUCCESS){
		return "Payment failed", errors.New(PAYMENT_FAILURE)
	}

	return "Payment successfully made",nil
}

func(ps PaymentStatusParams)GetUserPaymentStatus(userId string,accountId string)(models.UserPayment, error){

	var paymentObj models.UserPayment

	baseurl, errurl := url.Parse(paymentUrl)

	if errurl != nil {
		return models.UserPayment{}, errors.New("Malformed URL")
	}

	params := url.Values{}
	params.Add("userId",userId)
	params.Add("accountId",accountId )

	baseurl.RawQuery = params.Encode()

	resp, errResp := PaymentClientUtils.CallClient(ps.ClientInterface,"GET",baseurl.String(),nil)

	if errResp !=nil {
		return models.UserPayment{},errResp
	}

	err := json.Unmarshal(resp,&paymentObj)

	if err!= nil {
		return models.UserPayment{},err
	}

	return paymentObj,nil
}


