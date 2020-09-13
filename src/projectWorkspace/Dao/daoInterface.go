package Dao

import (
	"projectWorkspace/model"
)

type InterfaceDao interface{

	GetUser(string, *model.User)(error)

	
}