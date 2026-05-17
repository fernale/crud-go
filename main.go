package main

import (
	"context"
	"log"

	"github.com/fernale/crud-go/src/configuration/database/mongodb"
	"github.com/fernale/crud-go/src/configuration/logger"
	"github.com/fernale/crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	logger.Info("About to start the application")
	err:= godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	database, err := mongodb.NewMongoDBConnection(context.Background())

	if err != nil {
		logger.Error("Error trying to connect to database.", err)
		return
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
