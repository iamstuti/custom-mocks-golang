package MockDatastore

import (
	"cloud.google.com/go/datastore"
	

	"context"
	"errors"

	"github.com/googleapis/google-cloud-go-testing/datastore/dsiface"
)

type MockDatastoreClientErr struct {
	dsiface.Client
}

const (
	Err_PUT_Operation = "datastore:Error in executing of PUT operation"
	Err_Next          = "Error in returning the key of the next result"
)

var ErrConcurrentTransaction = errors.New("datastore:Error in executing of concurrent transaction")

//Implementing Failure i.e. when error occured during call
func (me *MockDatastoreClientErr) Close() error {
	return errors.New("Error Closing")
}
func (me *MockDatastoreClientErr) AllocateIDs(ctx context.Context, keys []*datastore.Key) ([]*datastore.Key, error) {
	return nil, errors.New("Error in Allocating IDs")
}
func (me *MockDatastoreClientErr) Count(ctx context.Context, q *datastore.Query) (n int, err error) {
	return 1, errors.New("Error in returning count")
}

func (me *MockDatastoreClientErr) Delete(ctx context.Context, key *datastore.Key) error {
	return errors.New("Error in Deletion")
}

func (me *MockDatastoreClientErr) DeleteMulti(ctx context.Context, keys []*datastore.Key) (err error) {
	return errors.New("Error in Deleting multiple entities")
}

func (me *MockDatastoreClientErr) Get(ctx context.Context, key *datastore.Key, dst interface{}) (err error) {
	return datastore.ErrNoSuchEntity
}

func (me *MockDatastoreClientErr) GetAll(ctx context.Context, q *datastore.Query, dst interface{}) (keys []*datastore.Key, err error) {
	return nil, errors.New("Error in fetching all keys")
}
func (me *MockDatastoreClientErr) GetMulti(ctx context.Context, keys []*datastore.Key, dst interface{}) (err error) {
	return errors.New("Error in fetching executing Get Multi function")
}

func (me *MockDatastoreClientErr) Mutate(ctx context.Context, muts ...*datastore.Mutation) (ret []*datastore.Key, err error) {
	return nil, errors.New("Error in fetching keys post Mutate operation")
}
func (me *MockDatastoreClientErr) NewTransaction(ctx context.Context, opts ...datastore.TransactionOption) (t dsiface.Transaction, err error) {
	return nil, errors.New("Error in running the new transaction")
}

func (me *MockDatastoreClientErr) Put(ctx context.Context, key *datastore.Key, src interface{}) (*datastore.Key, error) {
	return nil, errors.New(Err_PUT_Operation)
}

func (me *MockDatastoreClientErr) PutMulti(ctx context.Context, keys []*datastore.Key, src interface{}) (ret []*datastore.Key, err error) {
	return nil, errors.New("Error execution of Put operation on multiple keys")
}
func (me *MockDatastoreClientErr) Run(ctx context.Context, q *datastore.Query) dsiface.Iterator {
	return &MockIteratorPutError{err: errors.New(Err_Next)}
}

func (me *MockDatastoreClientErr) RunInTransaction(ctx context.Context, f func(tx dsiface.Transaction) error, opts ...datastore.TransactionOption) (cmt dsiface.Commit, err error) {
	var tx dsiface.Transaction
	if err := f(tx); err != ErrConcurrentTransaction {
		return &MockDSCommit{}, err

	}
	return &MockDSCommit{}, ErrConcurrentTransaction
}

type MockIteratorPutError struct {
	MockIterator
	err error
}

func (mte *MockIteratorPutError) Cursor() (datastore.Cursor, error) {
	return datastore.Cursor{}, errors.New("Error in returning a cursor for the iterator's current location")
}
func (mte *MockIteratorPutError) Next(dst interface{}) (*datastore.Key, error) {
	return nil, errors.New(Err_Next)
}