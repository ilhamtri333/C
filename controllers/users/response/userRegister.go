package response

import (
	"cleanarch/business/users"
	"time"

	"gorm.io/gorm"
)

type responseRegister struct {
	Id       uint           `json:"id"`
	CreateAt time.Time      `json:"createAt"`
	UpdateAt time.Time      `json:"updateAt"`
	DeleteAt gorm.DeletedAt `json:"deleteAt"`
	Nama     string         `json:"nama"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Token    string         `json:"token"`
}

func FromDomainToRegist(domain users.Domain) responseRegister {
	return responseRegister{
		Id:       domain.Id,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
		DeleteAt: domain.DeleteAt,
		Nama:     domain.Nama,
		Email:    domain.Email,
		Password: domain.Password,
		Token:    domain.Token,
	}
}
