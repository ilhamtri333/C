package owner

import (
	"cleanarch/drivers/database/kost"
	"context"
	"time"
)

type Domain struct {
	ID       uint
	Owner    string
	Email    string
	No_telp  int
	Kost     []kost.Kost
	CreateAt time.Time
	UpdateAt time.Time
}

type OwnerUseCaseInterface interface {
	InsertOwner(ctx context.Context, domain *Domain) (Domain, error)
	GetAllOwner(ctx context.Context) ([]Domain, error)
	GetOwnerById(ctx context.Context, id uint) (Domain, error)
	UpdateOwner(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteOwner(ctx context.Context, id uint) error
}

type OwnerRepoInterface interface {
	InsertOwner(ctx context.Context, domain *Domain) (Domain, error)
	GetAllOwner(ctx context.Context) ([]Domain, error)
	GetOwnerById(ctx context.Context, id uint) (Domain, error)
	UpdateOwner(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteOwner(ctx context.Context, id uint) error
}
