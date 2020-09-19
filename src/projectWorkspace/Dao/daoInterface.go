package Dao

import (
	"projectWorkspace/projectWorkspace/model"
)

type InterfaceDao interface{

	GetUser(string, *model.User)error

	GetAllAccounts()([]model.Account,error)

	AddUser(model.User)(string,error)


	
}