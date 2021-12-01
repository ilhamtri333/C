package response

import (
	payment "cleanarch/business/payment_method"
	"time"
)

type PaymentResponse struct {
	Id       uint      `json:"id"`
	Nama     string    `json:"nama"`
	Type     string    `json:"type"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type SearchPayment struct {
	Payment interface{}
}

func FromDomain(domain payment.Domain) PaymentResponse {
	return PaymentResponse{
		Id:       domain.Id,
		Nama:     domain.Nama,
		Type:     domain.Type,
		CreateAt: domain.CreatedAt,
		UpdateAt: domain.UpdatedAt,
	}
}

func PaymentAll(domain []payment.Domain) []PaymentResponse {
	var getAll []PaymentResponse
	for _, v := range domain {
		getAll = append(getAll, FromDomain(v))
	}
	return getAll
}
