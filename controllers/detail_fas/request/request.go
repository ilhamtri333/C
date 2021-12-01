package request

import detailfas "cleanarch/business/detail_fas"

type InsertDetailFas struct {
	KostID    uint `json:"kostid"`
	KostFasID uint `json:"kostfasid"`
}

func (df *InsertDetailFas) ToDomain() *detailfas.Domain {
	return &detailfas.Domain{
		KostID:    df.KostID,
		KostFasID: df.KostFasID,
	}
}
