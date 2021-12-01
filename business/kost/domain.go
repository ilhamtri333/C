package kost

import (
	kostfasilitas "cleanarch/drivers/database/kost_fasilitas"
	"context"
	"time"
)

type Domain struct {
	ID        uint
	Nama      string
	Type_Kost string
	Desc_Kost string
	Address   string
	Duration  int
	OwnerID   uint
	KostFas   kostfasilitas.KostFas
	CreateAt  time.Time
	UpdateAt  time.Time
}

type KostUseCaseInterface interface {
	InsertKost(ctx context.Context, domain *Domain) (Domain, error)
	GetAllKost(ctx context.Context) ([]Domain, error)
	GetKostById(ctx context.Context, id uint) (Domain, error)
	UpdateKost(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteKost(ctx context.Context, id uint) error
}

type KostRepoInterface interface {
	InsertKost(ctx context.Context, domain *Domain) (Domain, error)
	GetAllKost(ctx context.Context) ([]Domain, error)
	GetKostById(ctx context.Context, id uint) (Domain, error)
	UpdateKost(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteKost(ctx context.Context, id uint) error
}
