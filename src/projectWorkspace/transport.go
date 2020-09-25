package projectWorkspace

import (
	"encoding/json"
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"

)

var errBadRoute = errors.New("bad route")


func MakeHandler(ts IService, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	addUserHandler := kithttp.NewServer(
		AddUserEndpoint(ts),
		DecodeUserRequest,
		encodeResponse,
		opts...,
	)

	listUserHandler := kithttp.NewServer(
		GetUserEndpoint(ts),
		DecodeByIdRequest,
		encodeResponse,
		opts...,
	)

	allAccountsHandler := kithttp.NewServer(
		GetAccountListEndpoint(ts),
		DecodeAccountListRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/edu/v1/addUser", addUserHandler).Methods("POST")
	r.Handle("/edu/v1/getUser", listUserHandler).Methods("GET")
	r.Handle("/edu/v1/accounts", allAccountsHandler).Methods("GET")

	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func DecodeAccountListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return AllAcountsRequest{}, nil
}

func DecodeByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return GetByIdRequest{ID: id}, nil
}

func DecodeUserRequest(_ context.Context, r *http.Request)(interface{}, error){

	var userRequestObj AddUserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequestObj); err != nil {
		return nil, err
	}

	return userRequestObj,nil
}