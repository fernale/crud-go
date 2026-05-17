package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	EncryptPassword()

	GetID() string
	SetID(string)
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		id: "",
		email: email,
		password: password,
		name: name,
		age: age,
	}
}