package Dao

import(

	"log"
	"projectWorkspace/projectWorkspace/model"
	"projectWorkspace/projectWorkspace/GCP/Datastore"
	"cloud.google.com/go/datastore"
)

type Dao struct{
	InterfaceDao
	Datastore.GDatastore
}

func(dao Dao)GetUser(userId string, user *model.User)(error){

	key, errKey := Datastore.StringIDToDatastoreKey(userId, Datastore.GetUserKind())

	if errKey !=nil{
		log.Println(errKey.Error())
		return errKey
	}

	errGet := dao.IDatastoreClient.Get(dao.DbContext,key,user)

	if errGet != nil{
		log.Println(errGet.Error())
		return Datastore.ErrNoSuchEntity 
	}

	return nil
}

func (dao Dao)GetAllAccounts()([]model.Account){
	query := datastore.NewQuery(Datastore.GetAccountKind())
	var accountList  []model.Account

}



