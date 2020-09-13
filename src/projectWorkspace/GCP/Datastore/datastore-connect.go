package Datastore

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/googleapis/google-cloud-go-testing/datastore/dsiface"
)

//Datastore Errors
var (
	ErrInvalidEntityType = errors.New("datastore: invalid entity type")
	ErrInvalidKey        = errors.New("datastore : invalid key")
	ErrNoSuchEntity      = errors.New("datastore: no such entity")
)

//GCloudDatastore :
type GDatastore struct {
	DbContext context.Context
	IDatastoreClient  dsiface.Client
}

var namespace string

//Connect :
func (c *GDatastore) Connect(projectName string, nameSpace string) bool {

	c.DbContext = context.Background()
	client, err := datastore.NewClient(c.DbContext, projectName)

	if err != nil {
		log.Fatalf("Client creation failed: %v", err)
		return false
	}

	c.IDatastoreClient = dsiface.AdaptClient(client)

	if c.IDatastoreClient == nil {
		log.Println("Datastore Client initialization failed !!")
		return false
	}

	namespace = nameSpace
	log.Println("Datastore Client initialized successfully !!")

	return true
}
