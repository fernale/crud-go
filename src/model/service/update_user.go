package service

import (
	rest_err "github.com/fernale/crud-go/src/configuration/rest_err"
	"github.com/fernale/crud-go/src/model"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	return nil
}
//UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr