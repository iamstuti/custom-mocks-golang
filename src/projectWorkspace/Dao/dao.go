package Dao

import(

	"log"
	"fmt"
	"errors"
	"projectWorkspace/projectWorkspace/model"
	"projectWorkspace/projectWorkspace/GCP/Datastore"
	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
	
)

type Dao struct{
	InterfaceDao
	Datastore.GDatastore
}

const Err_EmptyKey="Datastore key cannot be empty"

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

func (dao Dao)AddUser(userObj model.User)(string, error)  {

	key := Datastore.GenerateIncompleteKey(Datastore.GetUserKind(),Datastore.GetNamespace(),nil)

	userKey, errAdd := dao.IDatastoreClient.Put(dao.DbContext,key,&userObj)

	if errAdd != nil{
		return "", errors.New("Error occurred in saving user object")

	}

	return fmt.Sprint(userKey.ID),nil
}

func (dao Dao)GetAllAccounts()([]model.Account,error){
	query := datastore.NewQuery(Datastore.GetAccountKind())
	queryNamespace:= query.Namespace(Datastore.GetNamespace())
	
	it := dao.IDatastoreClient.Run(dao.DbContext, queryNamespace)
	var accountList  []model.Account

	is_Done := false
	var ErrNext error
	var accountKey *datastore.Key

	for {

		var accObj model.Account


		accountKey, ErrNext = it.Next(&accObj)

		if ErrNext == iterator.Done{
			is_Done = true
			break
		}

		if ErrNext != nil {
			break
		}

		if accountKey == nil {
			ErrNext = errors.New(Err_EmptyKey)
			break
		}

		accObj.AccountId= fmt.Sprint(accountKey.ID)
		accountList = append(accountList,accObj)
	}




	if !is_Done{
		return []model.Account{},ErrNext
	}

	return accountList,nil
}



