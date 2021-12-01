package request

import payment "cleanarch/business/payment_method"

type InsertPayment struct {
	Type string `json:"type"`
	Nama string `json:"nama"`
}

func (pay *InsertPayment) ToDomain() *payment.Domain {
	return &payment.Domain{
		Type: pay.Type,
		Nama: pay.Nama,
	}
}
