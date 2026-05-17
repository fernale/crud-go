package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/fernale/crud-go/src/configuration/logger"
	"github.com/fernale/crud-go/src/model"
)


type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	EncryptPassword()
	GetJSONValue() (string, error)
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
		return &userDomain{
			email, password, name, age,
		}
}


type userDomain struct {
	Email string
	Password string
	Name string
	Age int8
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)

	if err != nil {
		logger.Error("error GetJSONValue marshal", err)
		return "", err
	}

	return string(b), nil
}

func (ud *userDomain) GetEmail()string {
	return ud.Email
}
func (ud *userDomain) GetPassword() string {
	return ud.Password
}
func (ud *userDomain) GetName() string {
	return ud.Name
}
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) EncryptPassword(){
	hash := sha256.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}


