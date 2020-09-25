package PaymentClientUtils

import (
	"io"
)

type ClientInterface interface{
	HTTPClientMethods(string,string,io.Reader)([]byte,error)
}
