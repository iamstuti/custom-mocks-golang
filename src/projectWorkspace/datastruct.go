package projectWorkspace

import (
	"projectWorkspace/projectWorkspace/model"
)


type AllAcountsRequest struct {

}

type GetByIdRequest struct{
	ID string
}

type AddUserRequest struct {
	UserName               string `json:"userName"`
	UserEMail              string `json:"userEmail"`
	AccountID              string  `json:"accountId"`
}

type CustomResponse struct {
	Response string `json:"Response"`
}

type GetUserResponse struct {
	User model.User `json:"UserDetails"`
}

type GetAccountListResponse struct {
	Account []model.Account `json:"AccountList"`
}