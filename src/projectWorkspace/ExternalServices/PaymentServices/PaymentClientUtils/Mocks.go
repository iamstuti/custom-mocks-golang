package PaymentClientUtils

import (
	"errors"
	"io"
	"encoding/json"
	"time"
	"projectWorkspace/projectWorkspace/ExternalServices/PaymentServices/models"
)

type MockClient struct {
	ClientInterface
}

const Err_Call = "Error occurred in calling client"

func (mockHTTPClient MockClient) HTTPClientMethods(method string, url string, body io.Reader) ([]byte, error) {


	var byteArr []byte
	
	switch method {
	case "GET":
		userObj := models.UserPayment{
			ReferenceNumber: 387460877,
			PaymentDate:     time.Now(),
			PaymentMode:     "HDFC Visa Credit Cark",
			Amount:          39.25,
			Status:          "success",
		}
	
		byteArr, _ = json.Marshal(userObj)
	
	case "POST":

		response := models.Response{
			Message:"Payment completed successfully",
		}

		byteArr,_ = json.Marshal(response)

	}

	return byteArr,nil

}

type MockClientError struct {
	ClientInterface
}

func (mockHTTPClientErr MockClientError) HTTPClientMethods(method string, url string, body io.Reader) ([]byte, error) {

	return []byte{}, errors.New(Err_Call)
}