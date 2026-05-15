package controller

import (
	"net/http"

	"github.com/fernale/crud-go/src/configuration/logger"
	"github.com/fernale/crud-go/src/configuration/validation"
	"github.com/fernale/crud-go/src/controller/model/request"
	"github.com/fernale/crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	
	response := response.UserResponse{
		ID: "test",
		Name: userRequest.Name,
		Email: userRequest.Email,
		Age: userRequest.Age,
	}

	logger.Info("User CreateUser successfully", zap.String("journey", "createUser"))
	c.JSON(http.StatusCreated, response)
}