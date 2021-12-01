package owner

import (
	"cleanarch/business/owner"
	"cleanarch/drivers/database/kost"
	"time"

	"gorm.io/gorm"
)

type Owner struct {
	ID        uint `gorm:"prim.aryKey"`
	Owner     string
	Email     string
	No_telp   int
	Kost      []kost.Kost
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

func (own Owner) ToDomain() owner.Domain {
	return owner.Domain{
		ID:       own.ID,
		Owner:    own.Owner,
		Email:    own.Email,
		No_telp:  own.No_telp,
		Kost:     own.Kost,
		CreateAt: own.CreatedAt,
		UpdateAt: own.UpdatedAt,
	}
}

func FromDomain(domain owner.Domain) Owner {
	return Owner{
		ID:        domain.ID,
		Owner:     domain.Owner,
		Email:     domain.Email,
		No_telp:   domain.No_telp,
		Kost:      domain.Kost,
		CreatedAt: domain.CreateAt,
		UpdatedAt: domain.UpdateAt,
	}
}

func AllOwner(dataown []Owner) []owner.Domain {
	All := []owner.Domain{}
	for _, v := range dataown {
		All = append(All, v.ToDomain())
	}
	return All
}
