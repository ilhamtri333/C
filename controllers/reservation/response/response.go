package response

import (
	"cleanarch/business/reservation"
	"time"
)

type ReservationResponse struct {
	Id               uint      `json:"id"`
	User_Id          uint      `json:"user_id"`
	Kost_Id          uint      `json:"kost_id"`
	PaymentMethod_Id uint      `json:"paymentmethod_id"`
	Check_In         string    `json:"check_in"`
	Check_Out        string    `json:"check_out"`
	CreateAt         time.Time `json:"createAt"`
	UpdateAt         time.Time `json:"updateAt"`
}

type SearchReservation struct {
	Reservation interface{}
}

func FromDomain(domain reservation.Domain) ReservationResponse {
	return ReservationResponse{
		Id:               domain.Id,
		User_Id:          domain.User_Id,
		Kost_Id:          domain.Kost_Id,
		PaymentMethod_Id: domain.PaymentMethod_Id,
		Check_In:         domain.Check_In,
		Check_Out:        domain.Check_Out,
		CreateAt:         domain.CreateAt,
		UpdateAt:         domain.UpdateAt,
	}
}

func ReservationAll(domain []reservation.Domain) []ReservationResponse {
	var getAll []ReservationResponse
	for _, v := range domain {
		getAll = append(getAll, FromDomain(v))
	}
	return getAll
}
