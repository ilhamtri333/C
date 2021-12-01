package request

import "cleanarch/business/owner"

type InsertOwner struct {
	Owner   string `json:"owner"`
	Email   string `json:"email"`
	No_telp int    `json:"no_telp"`
}

func (own *InsertOwner) ToDomain() *owner.Domain {
	return &owner.Domain{
		Owner:   own.Owner,
		Email:   own.Email,
		No_telp: own.No_telp,
	}
}
