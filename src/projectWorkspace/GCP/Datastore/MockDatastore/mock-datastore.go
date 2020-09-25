package MockDatastore

import (
	"projectWorkspace/projectWorkspace/model"
	"projectWorkspace/projectWorkspace/GCP/Datastore"
	"context"
	"fmt"
	"time"
	"cloud.google.com/go/datastore"
	"github.com/googleapis/google-cloud-go-testing/datastore/dsiface"
	"google.golang.org/api/iterator"
)


	 

type MockDatastoreClient struct {
	dsiface.Client
}

func GetMockDatastoreClient() Datastore.GDatastore {
	client := MockDatastoreClient{}

	return Datastore.GDatastore{IDatastoreClient: &client}
}

func GetMockDSClientPutError() Datastore.GDatastore {

	client := MockDatastoreClientErr{}

	return Datastore.GDatastore{IDatastoreClient: &client}
}

//Implementing Successful execution i.e. no error
func (mds *MockDatastoreClient) Close() error {
	return nil
}
func (mds *MockDatastoreClient) AllocateIDs(ctx context.Context, keys []*datastore.Key) ([]*datastore.Key, error) {
	return nil, nil
}
func (mds *MockDatastoreClient) Count(ctx context.Context, q *datastore.Query) (n int, err error) {
	return 1, nil
}

func (mds *MockDatastoreClient) Delete(ctx context.Context, key *datastore.Key) error {
	return nil
}

func (mds *MockDatastoreClient) DeleteMulti(ctx context.Context, keys []*datastore.Key) (err error) {
	return nil
}

func (mds *MockDatastoreClient) Get(ctx context.Context, key *datastore.Key, dst interface{}) (err error) {

	switch dst.(type) {

	case *model.User:

		user := dst.(*model.User)

		user.UserName = "Arianda Gary"
		user.UserEMail = "AGary12@gmail.com"
		user.AccountID = "AZ3129WR2M"
		user.UserId = "5766313779658752"
		user.SubscriptionID = "ZCI6Ik1UVTRNRE0xT0RrMk9BPT0ifQ"
		user.PaymentDate= time.Now().AddDate(0,-3,3)
		user.SubscriptionExpiryDate = time.Now().AddDate(1,-3,2)
		user.PaymentMethod ="visa credit card"



	case *model.Account:

		account := dst.(*model.Account)

		account.AccountId = "AZ3129WR2M"
		account.AccountName = "Amazon US Account"
		account.BillingID = "US9011B2"

		var svcObj model.Services
		svcObj.ServiceName = "Kindle Unlimited Subscription"
		svcObj.ServiceMonthlyCost = "3"
		svcObj.ServiceId ="KMQAZ312"
		svcObj.AnnualCost= "28"
		var svcObjList []model.Services
		svcObjList= append(svcObjList,svcObj)
		account.ServiceList=svcObjList

		
		

	}

	return nil

}

func (mds *MockDatastoreClient) GetAll(ctx context.Context, q *datastore.Query, dst interface{}) (keys []*datastore.Key, err error) {
	return nil, nil
}
func (mds *MockDatastoreClient) GetMulti(ctx context.Context, keys []*datastore.Key, dst interface{}) (err error) {
	return nil
}

func (mds *MockDatastoreClient) Mutate(ctx context.Context, muts ...*datastore.Mutation) (ret []*datastore.Key, err error) {
	return nil, nil
}
func (mds *MockDatastoreClient) NewTransaction(ctx context.Context, opts ...datastore.TransactionOption) (t dsiface.Transaction, err error) {
	return nil, nil
}

func (mds *MockDatastoreClient) Put(ctx context.Context, key *datastore.Key, src interface{}) (*datastore.Key, error) {
	return &datastore.Key{ID: 1234}, nil
}

func (mds *MockDatastoreClient) PutMulti(ctx context.Context, keys []*datastore.Key, src interface{}) (ret []*datastore.Key, err error) {
	return nil, nil
}

func (mds *MockDatastoreClient) Run(ctx context.Context, q *datastore.Query) dsiface.Iterator {



	var userList []model.User

	userObj1 := model.User{UserName: "Arianda Gary",
	UserId:        "5766313779696200",
	UserEMail:           "AGary12@gmail.com",
	AccountID:         "AZ8D7DK",
	PaymentDate:   time.Now().AddDate(-1,1,1),
	SubscriptionExpiryDate: time.Now().AddDate(0,1,0),
	SubscriptionID:"ZCI6Ik1UVXEQ420xT0RrMk9BPT0ifQ"}

	userObj2 := model.User{UserName: "Zack Wellington",
	UserId:        "5766313779658752",
	UserEMail:           "ZackWG@gmail.com",
	AccountID:         "AZ3129WR2M",
	PaymentDate:   time.Now().AddDate(-1,2,3),
	SubscriptionExpiryDate: time.Now().AddDate(0,2,2),
	SubscriptionID:"ZMEX4Ik1UVXEQ420xT0RsBk4BPT0ifQ"}

	userList = append(userList, userObj1)
	userList = append(userList, userObj2)

	var svcObj model.Services
	var svcObj2 model.Services
	var svcObjList []model.Services


	svcObj.ServiceName = "Kindle Unlimited Subscription"
	svcObj.ServiceMonthlyCost = "3"
	svcObj.AnnualCost= "28"
	svcObj.ServiceId ="KMQAZ312"

	svcObj2.ServiceName = "Audible Subscription"
	svcObj2.ServiceMonthlyCost = "6"
	svcObj2.AnnualCost= "64"
	svcObj.ServiceId ="ALWAZ312"


	svcObjList= append(svcObjList,svcObj)
	svcObjList= append(svcObjList,svcObj2)


	var account model.Account
	var accountList []model.Account

	account.AccountId = "AZ3129WR2M"
	account.AccountName = "Amazon US Account"
	account.BillingID = "US9011B2"
	account.ServiceList=svcObjList

	
	mIterator := MockIterator{userList: userList, accounts: accountList,services:svcObjList,Done: iterator.Done}
	return &mIterator
}

func (mds *MockDatastoreClient) RunInTransaction(ctx context.Context, f func(tx dsiface.Transaction) error, opts ...datastore.TransactionOption) (cmt dsiface.Commit, err error) {
	var tx dsiface.Transaction
	if err := f(tx); err != nil {
		return nil, err

	}

	return &MockDSCommit{}, nil
}

type MockIterator struct {
	dsiface.Iterator
	userList []model.User
	accounts []model.Account
	services []model.Services
	currentIdx   int
	Done 		 error
}

func (mt *MockIterator) Cursor() (datastore.Cursor, error) { return datastore.Cursor{}, nil }

func (mt *MockIterator) Next(dst interface{}) (*datastore.Key, error) {
	switch dst := dst.(type) {
	case *model.Account:

		if mt.currentIdx < len(mt.accounts) {

			accountList := mt.accounts[mt.currentIdx]

			key := Datastore.GenerateDataStoreNameKey(Datastore.GetAccountKind(), accountList.AccountId, nil)

			obj := dst
			obj.AccountId = accountList.AccountId
			obj.AccountName = accountList.AccountName
			obj.BillingID = accountList.BillingID
			obj.ServiceList = accountList.ServiceList
			

			mt.currentIdx = mt.currentIdx + 1

			return key, nil

		}
		return nil, mt.Done

	case *model.Services:
		if mt.currentIdx < len(mt.services) {

			infoListObj := mt.services[mt.currentIdx]
			key := Datastore.GenerateDataStoreNameKey(Datastore.GetServicesKind(),infoListObj.ServiceId,nil)

			obj := dst
			obj.ServiceName = infoListObj.ServiceName
			obj.ServiceMonthlyCost = infoListObj.ServiceMonthlyCost
			obj.ServiceId = infoListObj.ServiceId
			

			mt.currentIdx = mt.currentIdx + 1
			return key, nil

		}
		return nil, mt.Done

	case *model.User:

		if mt.currentIdx < len(mt.userList) {

			uList := mt.userList[mt.currentIdx]
			key, _ := Datastore.StringIDToDatastoreKey(uList.UserId, Datastore.GetUserKind())

			obj := dst
			obj.UserName = uList.UserName
			obj.UserId = uList.UserId
			obj.AccountID = uList.AccountID
			obj.PaymentDate = uList.PaymentDate
			obj.SubscriptionExpiryDate = uList.SubscriptionExpiryDate
			obj.SubscriptionID = uList.SubscriptionID

			mt.currentIdx = mt.currentIdx + 1
			return key, nil

		}
		return nil,  mt.Done


	default:
		fmt.Printf("The type of v is unknown\n")
	}

	return nil, nil

}

type MockDSTransaction struct {
	dsiface.Transaction
}

type MockDSCommit struct {
	dsiface.Commit
}

func (mt *MockDSTransaction) Commit() (c dsiface.Commit, err error) {
	return &MockDSCommit{}, nil
}

func (mdc *MockDSCommit) Key(p *datastore.PendingKey) *datastore.Key {
	return nil
}
