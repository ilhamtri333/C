package request

import "cleanarch/business/reservation"

type InsertReservation struct {
	User_Id          uint   `json:"user_id"`
	Kost_Id          uint   `json:"kost_id"`
	PaymentMethod_Id uint   `json:"paymentmethod_id"`
	Check_In         string `json:"check_in"`
	Check_Out        string `json:"check_out"`
}

func (rs *InsertReservation) ToDomain() *reservation.Domain {
	return &reservation.Domain{
		User_Id:          rs.User_Id,
		Kost_Id:          rs.Kost_Id,
		PaymentMethod_Id: rs.PaymentMethod_Id,
		Check_In:         rs.Check_In,
		Check_Out:        rs.Check_Out,
	}
}
