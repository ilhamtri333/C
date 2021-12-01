package detailfas

import (
	detailfas "cleanarch/business/detail_fas"
	"time"

	"gorm.io/gorm"
)

type DetailFas struct {
	ID        uint `gorm:"prim.aryKey"`
	KostID    uint
	KostFasID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

func (df DetailFas) ToDomain() detailfas.Domain {
	return detailfas.Domain{
		ID:        df.ID,
		KostID:    df.KostID,
		KostFasID: df.KostFasID,
		CreateAt:  df.CreatedAt,
		UpdateAt:  df.UpdatedAt,
	}
}

func FromDomain(domain detailfas.Domain) DetailFas {
	return DetailFas{
		ID:        domain.ID,
		KostID:    domain.KostID,
		KostFasID: domain.KostFasID,
		CreatedAt: domain.CreateAt,
		UpdatedAt: domain.UpdateAt,
	}
}

func AllDetailFas(datadf []DetailFas) []detailfas.Domain {
	All := []detailfas.Domain{}
	for _, v := range datadf {
		All = append(All, v.ToDomain())
	}
	return All
}
