package payment

import (
	payment "cleanarch/business/payment_method"
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	Id        uint `gorm:"primaryKey"`
	Type      string
	Nama      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

func (pay Payment) ToDomain() payment.Domain {
	return payment.Domain{
		Id:        pay.Id,
		Type:      pay.Type,
		Nama:      pay.Nama,
		CreatedAt: pay.CreatedAt,
		UpdatedAt: pay.UpdatedAt,
	}
}

func FromDomain(domain payment.Domain) Payment {
	return Payment{
		Id:        domain.Id,
		Type:      domain.Type,
		Nama:      domain.Nama,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func AllPayment(datapayment []Payment) []payment.Domain {
	All := []payment.Domain{}
	for _, v := range datapayment {
		All = append(All, v.ToDomain())
	}
	return All
}
