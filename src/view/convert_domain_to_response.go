package view

import (
	"github.com/fernale/crud-go/src/controller/model/response"
	"github.com/fernale/crud-go/src/model"
)


func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse{
	return response.UserResponse{
		ID: userDomain.GetID(),
		Name: userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age: userDomain.GetAge(),
	}
}