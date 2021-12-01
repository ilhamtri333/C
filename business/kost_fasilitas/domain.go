package kostfasilitas

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	Fasilitas string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type KostFasUsecaseInterface interface {
	InsertKostFas(ctx context.Context, domain *Domain) (Domain, error)
	GetAllKostFas(ctx context.Context) ([]Domain, error)
	GetKostFasById(ctx context.Context, id uint) (Domain, error)
	UpdateKostFas(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteKostFas(ctx context.Context, id uint) error
}

type KostFasMetodRepoInterface interface {
	InsertKostFas(ctx context.Context, domain *Domain) (Domain, error)
	GetAllKostFas(ctx context.Context) ([]Domain, error)
	GetKostFasById(ctx context.Context, id uint) (Domain, error)
	UpdateKostFas(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteKostFas(ctx context.Context, id uint) error
}
