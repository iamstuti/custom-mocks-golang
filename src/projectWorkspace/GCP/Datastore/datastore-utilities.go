package Datastore

import (
	"context"
	"log"
	"strconv"

	"cloud.google.com/go/datastore"
	"github.com/googleapis/google-cloud-go-testing/datastore/dsiface"
)

func DatastorePut(Client dsiface.Client, parentKey *datastore.Key, kind string, data interface{}) (*datastore.Key, error) {

	context := context.Background()

	key := datastore.IncompleteKey(kind, parentKey)

	key.Namespace = namespace

	genKey, errKey := Client.Put(context, key, data)

	if errKey != nil {
		log.Println(errKey)
	}

	return genKey, errKey
}

func DatastorePutWithKey(Client dsiface.Client, selfKey *datastore.Key, parentKey *datastore.Key, kind string, data interface{}) (*datastore.Key, error) {

	context := context.Background()

	genKey, errKey := Client.Put(context, selfKey, data)

	if errKey != nil {
		log.Printf("DatastorePutWithKey error %s", errKey.Error())
	}

	return genKey, errKey
}

func GenerateDataStoreNameKey(kind string, name string, parent *datastore.Key) *datastore.Key {
	return datastore.NameKey(kind, name, parent)
}

func GenerateIncompleteKey(kind string, namespace string, parent *datastore.Key) *datastore.Key {
	key := datastore.IncompleteKey(kind, parent)
	key.Namespace = namespace
	return key
}

func StringIDToDatastoreKey(id string, kind string) (*datastore.Key, error) {
	ID, errstrconv := strconv.ParseInt(id, 10, 64)

	if errstrconv != nil {
		log.Printf("generateDatastoreKeyFromStringID() strtoint() %v", errstrconv.Error())
		return nil, errstrconv
	}

	key := datastore.IDKey(kind, ID, nil)
	key.Namespace = namespace 

	return key, errstrconv
}