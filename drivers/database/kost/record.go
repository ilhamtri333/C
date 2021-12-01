package kost

import (
	"cleanarch/business/kost"
	"time"

	"gorm.io/gorm"
)

type Kost struct {
	ID        uint `gorm:"primaryKey"`
	Nama      string
	Type_Kost string
	Desc_Kost string
	Address   string
	Duration  int
	OwnerID   uint
	CreateAt  time.Time
	UpdateAt  time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

func (kos Kost) ToDomain() kost.Domain {
	return kost.Domain{
		ID:        kos.ID,
		Nama:      kos.Nama,
		Type_Kost: kos.Type_Kost,
		Desc_Kost: kos.Desc_Kost,
		Address:   kos.Address,
		Duration:  kos.Duration,
		OwnerID:   kos.OwnerID,
		CreateAt:  kos.CreateAt,
		UpdateAt:  kos.UpdateAt,
	}
}

func FromDomain(domain kost.Domain) Kost {
	return Kost{
		ID:        domain.ID,
		Nama:      domain.Nama,
		Type_Kost: domain.Type_Kost,
		Desc_Kost: domain.Desc_Kost,
		Address:   domain.Address,
		Duration:  domain.Duration,
		OwnerID:   domain.OwnerID,
		CreateAt:  domain.CreateAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func AllKost(datakost []Kost) []kost.Domain {
	All := []kost.Domain{}
	for _, v := range datakost {
		All = append(All, v.ToDomain())
	}
	return All
}
