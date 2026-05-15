package controller

import (
	"fmt"
	"net/http"

	"github.com/fernale/crud-go/src/configuration/validation"
	"github.com/fernale/crud-go/src/controller/model/request"
	"github.com/fernale/crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	fmt.Println(userRequest)
	
	response := response.UserResponse{
		ID: "test",
		Name: userRequest.Name,
		Email: userRequest.Email,
		Age: userRequest.Age,
	}
	c.JSON(http.StatusCreated, response)
}