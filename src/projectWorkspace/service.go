package projectWorkspace

import (
	"log"
	"projectWorkspace/projectWorkspace/GCP/Datastore"
	"errors"
	"projectWorkspace/projectWorkspace/model"
	sendgrid "projectWorkspace/projectWorkspace/SendgridMail"
	dao  "projectWorkspace/projectWorkspace/Dao"
	PaymentInterfaces "projectWorkspace/projectWorkspace/ExternalServices/PaymentServices/InterfacesAndMocks"
	paymentSvc "projectWorkspace/projectWorkspace/ExternalServices/PaymentServices"
	httpUtils "projectWorkspace/projectWorkspace/ExternalServices/PaymentServices/PaymentClientUtils"

)

type IService interface{
	GetUser(string)(model.User,error)
	GetAllAccounts()([]model.Account,error)
	AddUser(AddUserRequest)(string,error)

	InitializeService(datastoreObj Datastore.GDatastore)(bool)
}

func SendMail(userobj model.User,client sendgrid.SendInterface)error{
	return sendgrid.TriggerMail(userobj,client) 
}


type Service struct{
	IDAO dao.InterfaceDao
	IMailClient sendgrid.SendInterface
	IPaymentInterface PaymentInterfaces.CheckPaymentInterface
}

var ErrInvalidArgument = errors.New("Invalid argument")

func (svc *Service)InitializeService(datastoreObj Datastore.GDatastore) bool{

	svc.IDAO = dao.Dao{GDatastore: datastoreObj}

	actualClient := httpUtils.HTTPClient{}

	paymentParamsObj := paymentSvc.PaymentStatusParams{ClientInterface: actualClient}

	svc.IPaymentInterface = paymentSvc.CheckPaymentStatus{PaymentStatusParams:paymentParamsObj}

	return true
}

func (svc Service)AddUser(userObj AddUserRequest)(string,error){
	
	var newUserObj model.User

	newUserObj.UserName = userObj.UserName
	newUserObj.UserEMail = userObj.UserEMail
	newUserObj.AccountID = userObj.AccountID
	
	msg,paymentInfo, errPayment := svc.IPaymentInterface.CheckPaymentStatus(userObj.UserEMail,userObj.AccountID)

	log.Printf("Payment Status : %v",msg)
	
	if errPayment != nil {
		return "",errors.New("Payment is pending")
	}

	newUserObj.PaymentDate=paymentInfo.PaymentDate
	newUserObj.PaymentMethod=paymentInfo.PaymentMode

	newUserObj.SubscriptionExpiryDate = newUserObj.PaymentDate.AddDate(1,0,0)

	_,errAdd := svc.IDAO.AddUser(newUserObj)

	if errAdd != nil{
		return "", errors.New("Error occurred in saving User")
	}

	
	return "Successfully added user",nil

}

func(svc Service)GetUser(userId string)(model.User,error){

	if userId == ""{
		return model.User{},ErrInvalidArgument
	}

	var userObj model.User 

	errGet:= svc.IDAO.GetUser(userId, &userObj)

	if errGet !=nil {
		return model.User{},errors.New("Error occurred in getting user with user Id: "+userId)
	}

	return userObj,nil
}

func(svc Service)GetAllAccounts()([]model.Account,error){

	accList, errList := svc.IDAO.GetAllAccounts()

	if errList !=nil {
		return []model.Account{}, errors.New("Error occurred in getting account list")
	}
	
	return accList, nil

}

