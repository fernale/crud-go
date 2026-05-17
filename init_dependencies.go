package main

import (
	"github.com/fernale/crud-go/src/controller"
	"github.com/fernale/crud-go/src/model/repository"
	"github.com/fernale/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	//init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return  controller.NewUserControllerInterface(service)
}