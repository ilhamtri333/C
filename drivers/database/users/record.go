package users

import (
	"cleanarch/business/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id       uint           `gorm:"primaryKey"`
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
	Nama     string
	Email    string
	Password string
}

func (user User) ToDomain() users.Domain {
	return users.Domain{
		Id:       user.Id,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
		DeleteAt: user.DeleteAt,
		Nama:     user.Nama,
		Email:    user.Email,
		Password: user.Password,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:       domain.Id,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
		DeleteAt: domain.DeleteAt,
		Nama:     domain.Nama,
		Email:    domain.Email,
		Password: domain.Password,
	}
}

func ListUserToDomain(data []User) []users.Domain {
	All := []users.Domain{}
	for _, v := range data {
		All = append(All, v.ToDomain())
	}
	return All
}
