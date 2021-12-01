package response

import (
	"cleanarch/business/users"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id       uint           `json:"id"`
	CreateAt time.Time      `json:"createAt"`
	UpdateAt time.Time      `json:"updateAt"`
	DeleteAt gorm.DeletedAt `json:"deleteAt"`
	Nama     string         `json:"nama"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Token    string         `json:"token"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
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
func FromListDomain(domain []users.Domain) []UserResponse {
	var All []UserResponse
	for _, v := range domain {
		All = append(All, FromDomain(v))
	}
	return All
}
