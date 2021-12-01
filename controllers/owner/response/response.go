package response

import (
	"cleanarch/business/owner"
	"cleanarch/drivers/database/kost"
	"time"
)

type OwnerResponse struct {
	ID       uint        `json:"id"`
	Owner    string      `json:"owner"`
	Email    string      `json:"email"`
	No_telp  int         `json:"no_telp"`
	Kost     []kost.Kost `json:"kost"`
	CreateAt time.Time   `json:"createAt"`
	UpdateAt time.Time   `json:"updateAt"`
}

type SearchOwner struct {
	Owner interface{}
}

func FromDomain(domain owner.Domain) OwnerResponse {
	return OwnerResponse{
		ID:       domain.ID,
		Owner:    domain.Owner,
		Email:    domain.Email,
		No_telp:  domain.No_telp,
		Kost:     domain.Kost,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
	}
}

func OwnerAll(domain []owner.Domain) []OwnerResponse {
	var getAll []OwnerResponse
	for _, v := range domain {
		getAll = append(getAll, FromDomain(v))
	}
	return getAll
}
