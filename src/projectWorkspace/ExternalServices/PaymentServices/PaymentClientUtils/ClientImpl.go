package PaymentClientUtils

import (
	"io/ioutil"
	"log"
	"io"
	"net/http"
)

type HTTPClient struct {
	ClientInterface
}

func (httpClient HTTPClient) HTTPClientMethods(method string, url string, body io.Reader) ([]byte, error) {

	client := &http.Client{}

	var responseBody []byte
	var req *http.Request
	var errReq error

	switch method {
	case "GET":
		req,errReq = http.NewRequest("GET",url,nil)


	
	case "POST":

		req,errReq = http.NewRequest("POST",url,body)

	}

	if errReq != nil {
		log.Printf("Error: %v",errReq.Error())
		return []byte{},errReq
	}

	resp, errResp := client.Do(req)

	if errResp != nil {
		log.Printf("Error: %v",errResp.Error())
		return []byte{},errResp
	}

	responseBody, errBody := ioutil.ReadAll(resp.Body)

	if errBody != nil {
		return []byte{},errBody
	}

	return responseBody,nil


}

func CallClient(clientObj ClientInterface, method string,url string,body io.Reader)([]byte,error){

	return clientObj.HTTPClientMethods(method,url,body)
}
