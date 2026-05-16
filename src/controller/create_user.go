package controller

import (
	"net/http"

	"github.com/fernale/crud-go/src/configuration/logger"
	"github.com/fernale/crud-go/src/configuration/validation"
	"github.com/fernale/crud-go/src/controller/model/request"
	"github.com/fernale/crud-go/src/model"
	"github.com/fernale/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context){
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age)
	
	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err !=nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User CreateUser successfully", zap.String("journey", "createUser"))
	c.String(http.StatusCreated, "")
}