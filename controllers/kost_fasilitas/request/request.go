package request

import kostfasilitas "cleanarch/business/kost_fasilitas"

type InsertKostFas struct {
	Fasilitas string `json:"fasilitas"`
}

func (fas *InsertKostFas) ToDomain() *kostfasilitas.Domain {
	return &kostfasilitas.Domain{
		Fasilitas: fas.Fasilitas,
	}
}
