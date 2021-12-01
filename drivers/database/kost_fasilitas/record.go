package kostfasilitas

import (
	kostfasilitas "cleanarch/business/kost_fasilitas"
	"time"

	"gorm.io/gorm"
)

type KostFas struct {
	Id        uint `gorm:"primaryKey"`
	Fasilitas string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

func (fas KostFas) ToDomain() kostfasilitas.Domain {
	return kostfasilitas.Domain{
		Id:        fas.Id,
		Fasilitas: fas.Fasilitas,
		CreatedAt: fas.CreatedAt,
		UpdatedAt: fas.UpdatedAt,
	}
}

func FromDomain(domain kostfasilitas.Domain) KostFas {
	return KostFas{
		Id:        domain.Id,
		Fasilitas: domain.Fasilitas,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func AllKostFas(datafas []KostFas) []kostfasilitas.Domain {
	All := []kostfasilitas.Domain{}
	for _, v := range datafas {
		All = append(All, v.ToDomain())
	}
	return All
}
