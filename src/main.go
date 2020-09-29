package main

import(
	"projectWorkspace/projectWorkspace"
	"projectWorkspace/projectWorkspace/GCP/Datastore"
	sendgrid "projectWorkspace/projectWorkspace/SendgridMail"
	goKitLogger "github.com/go-kit/kit/log"
	"net/http"
	"os"
	"log"



)

var dataStoreClient Datastore.GDatastore
var mailClient sendgrid.MailingClient

func InitService () bool {

	if !dataStoreClient.Connect("YOUR-PROJECT-ID",Datastore.GetNamespace()){
		log.Println("Datastore connection failed to establish")
		return false
	}

	if !mailClient.Connect("YOUR-Sendgrid-Key"){
		log.Println("Sendgrid connection failed to establish")
		return false
	}

	return true

}

func main(){

	if !InitService(){
		log.Println("Service initialization failed")
		return
	}

	var logger goKitLogger.Logger
	logger = goKitLogger.NewLogfmtLogger(goKitLogger.NewSyncWriter(os.Stderr))
	logger = goKitLogger.With(logger, "ts", goKitLogger.DefaultTimestampUTC)

	
	var ITestService projectWorkspace.IService

	serviceObj:= projectWorkspace.Service{}
	
	if serviceObj.InitializeService(dataStoreClient){
		log.Println("Service initialization successful")
	}

	ITestService = &serviceObj

	mux := http.NewServeMux()
	httpLogger := goKitLogger.With(logger, "component", "http")


	mux.Handle("/testService/v1/", projectWorkspace.MakeHandler(ITestService, httpLogger))
	http.ListenAndServe(":8080", nil)
}