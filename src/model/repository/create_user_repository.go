package repository

import (
	"context"
	"os"

	"github.com/fernale/crud-go/src/configuration/logger"
	rest_err "github.com/fernale/crud-go/src/configuration/rest_err"
	"github.com/fernale/crud-go/src/model"
)

var (
	MONGODB_USER_COLLECTION="MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userDomain.GetJSONValue();

	collection.InsertOne(context.Background())
}