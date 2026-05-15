package model

import (
	"fmt"

	"github.com/fernale/crud-go/src/configuration/logger"
	rest_err "github.com/fernale/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()
	fmt.Println(ud.Password)
	return nil
}