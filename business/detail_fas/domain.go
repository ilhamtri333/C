package detailfas

import (
	"context"
	"time"
)

type Domain struct {
	ID     uint
	KostID uint
	// Kost       kost.Domain
	KostFasID uint
	// KostFas    kostfasilitas.Domain
	CreateAt time.Time
	UpdateAt time.Time
}

type DetailFasUseCaseInterface interface {
	InsertDetailFas(ctx context.Context, domain *Domain) (Domain, error)
	GetAllDetailFas(ctx context.Context) ([]Domain, error)
	GetDetailFasById(ctx context.Context, id uint) (Domain, error)
	UpdateDetailFas(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteDetailFas(ctx context.Context, id uint) error
}

type DetailFasRepoInterface interface {
	InsertDetailFas(ctx context.Context, domain *Domain) (Domain, error)
	GetAllDetailFas(ctx context.Context) ([]Domain, error)
	GetDetailFasById(ctx context.Context, id uint) (Domain, error)
	UpdateDetailFas(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteDetailFas(ctx context.Context, id uint) error
}
