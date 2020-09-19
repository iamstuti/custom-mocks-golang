package projectWorkspace

import (
	"errors"
	"projectWorkspace/projectWorkspace/model"
	sendgrid "projectWorkspace/projectWorkspace/SendgridMail"
	dao  "projectWorkspace/projectWorkspace/Dao"
)

type IService interface{
	GetUser(string)(model.User,error)
	GetAllAccounts()([]model.Account,error)
	AddUser(model.User)(string,error)
}

func SendMail(model.User,sendgrid.SendInterface)error{
	return nil
}


type Service struct{
	IDAO dao.InterfaceDao
	IMailClient sendgrid.SendInterface
}

func (svc Service)AddUser(userObj model.User)(string,error){
	
	_,errAdd := svc.IDAO.AddUser(userObj)

	if errAdd != nil{
		return "", errors.New("Error occurred in saving User")
	}

	errSend := SendMail(userObj,svc.IMailClient)

	if errSend !=nil{
		return "", errors.New("Error occured in sending mail")
	}

	return "Successfully added user",nil

}

func(svc Service)GetUser(userId string)(model.User,error){

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

