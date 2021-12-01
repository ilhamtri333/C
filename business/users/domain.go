package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id       uint
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt gorm.DeletedAt `gorm:"index"`
	Nama     string
	Email    string
	Password string
	Token    string
}

type UserUseCaseInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	GetAllUser(ctx context.Context) ([]Domain, error)
	Register(domain *Domain, ctx context.Context) (Domain, error)
	GetUserById(ctx context.Context, id uint) (Domain, error)
	UpdateUser(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteUser(ctx context.Context, id uint) error
}

type UserRepoInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	GetAllUser(ctx context.Context) ([]Domain, error)
	Register(domain *Domain, ctx context.Context) (Domain, error)
	GetUserById(ctx context.Context, id uint) (Domain, error)
	UpdateUser(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteUser(ctx context.Context, id uint) error
}
