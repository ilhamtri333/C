package response

import (
	detailfas "cleanarch/business/detail_fas"
	"time"
)

type DetailFasResponse struct {
	ID        uint      `json:"id"`
	KostID    uint      `json:"kostid"`
	KostFasID uint      `json:"kostfasid"`
	CreateAt  time.Time `json:"createAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

func FromDomain(domain detailfas.Domain) DetailFasResponse {
	return DetailFasResponse{
		ID:        domain.ID,
		KostID:    domain.KostID,
		KostFasID: domain.KostFasID,
		CreateAt:  domain.CreateAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func DetailFasAll(domain []detailfas.Domain) []DetailFasResponse {
	var getAll []DetailFasResponse
	for _, v := range domain {
		getAll = append(getAll, FromDomain(v))
	}
	return getAll
}
