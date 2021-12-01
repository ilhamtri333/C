package request

import "cleanarch/business/kost"

type InsertKost struct {
	Nama      string `json:"nama"`
	Type_Kost string `json:"type_kost"`
	Desc_Kost string `json:"desc_kost"`
	Address   string `json:"address"`
	Duration  int    `json:"duration"`
	OwnerID   uint   `json:"ownerID"`
	// Owner     owner.Domain
}

func (kos *InsertKost) ToDomain() *kost.Domain {
	return &kost.Domain{
		Nama:      kos.Nama,
		Type_Kost: kos.Address,
		Desc_Kost: kos.Desc_Kost,
		Address:   kos.Address,
		Duration:  kos.Duration,
		OwnerID:   kos.OwnerID,
	}
}
