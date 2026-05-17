package main

import (
	"log"

	"github.com/fernale/crud-go/src/configuration/database/mongodb"
	"github.com/fernale/crud-go/src/configuration/logger"
	"github.com/fernale/crud-go/src/controller"
	"github.com/fernale/crud-go/src/controller/routes"
	"github.com/fernale/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	logger.Info("About to start the application")
	err:= godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	mongodb.NewMongoDBConnection()

	//init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
