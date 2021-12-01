package response

import (
	"cleanarch/business/kost"
	"time"
)

type KostResponse struct {
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	Type_Kost string `json:"type_kost"`
	Desc_Kost string `json:"desc_kost"`
	Address   string `json:"address"`
	Duration  int    `json:"duration"`
	OwnerID   uint   `json:"ownerid"`
	// Owner     owner.Domain
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func FromDomain(domain kost.Domain) KostResponse {
	return KostResponse{
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

func KostAll(domain []kost.Domain) []KostResponse {
	var getAll []KostResponse
	for _, v := range domain {
		getAll = append(getAll, FromDomain(v))
	}
	return getAll
}
