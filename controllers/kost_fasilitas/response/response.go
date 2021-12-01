package response

import (
	kostfasilitas "cleanarch/business/kost_fasilitas"
	"time"
)

type KostFasResponse struct {
	Id        uint      `json:"id"`
	Fasilitas string    `json:"fasilitas"`
	CreateAt  time.Time `json:"createAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

type SearchKostFas struct {
	KostFas interface{}
}

func FromDomain(domain kostfasilitas.Domain) KostFasResponse {
	return KostFasResponse{
		Id:        domain.Id,
		Fasilitas: domain.Fasilitas,
		CreateAt:  domain.CreatedAt,
		UpdateAt:  domain.UpdatedAt,
	}
}

func KostFasAll(domain []kostfasilitas.Domain) []KostFasResponse {
	var getAll []KostFasResponse
	for _, v := range domain {
		getAll = append(getAll, FromDomain(v))
	}
	return getAll
}
