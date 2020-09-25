package projectWorkspace

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func AddUserEndpoint(svc IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddUserRequest)

		msg,err := svc.AddUser(req)

		if err !=nil {
			return CustomResponse{},err
		}

		return CustomResponse{Response: msg}, nil
	}
}

func GetUserEndpoint(svc IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIdRequest)

		user,err := svc.GetUser(req.ID)

		if err !=nil {
			return GetUserResponse{},err
		}

		return GetUserResponse{User: user}, nil
	}
}

func GetAccountListEndpoint(svc IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		list,err := svc.GetAllAccounts()

		if err !=nil {
			return GetAccountListResponse{},err
		}

		return GetAccountListResponse{Account: list}, nil
	}
}