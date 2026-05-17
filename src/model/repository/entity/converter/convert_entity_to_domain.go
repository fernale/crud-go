package converter

import (
	"github.com/fernale/crud-go/src/model"
	"github.com/fernale/crud-go/src/model/repository/entity"
)

func ConvertEntityToDomain(en entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(en.Email, en.Password, en.Name, en.Age)
	domain.SetID(en.ID.Hex())
	return domain
}