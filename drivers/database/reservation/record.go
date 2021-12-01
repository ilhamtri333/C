package reservation

import (
	"cleanarch/business/reservation"
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	Id               uint `gorm:"primaryKey"`
	User_Id          uint
	Kost_Id          uint
	PaymentMethod_Id uint
	Check_In         string
	Check_Out        string
	CreateAt         time.Time
	UpdateAt         time.Time
	DeleteAt         gorm.DeletedAt `gorm:"index"`
}

func (rs Reservation) ToDomain() reservation.Domain {
	return reservation.Domain{
		Id:               rs.Id,
		User_Id:          rs.User_Id,
		Kost_Id:          rs.Kost_Id,
		PaymentMethod_Id: rs.PaymentMethod_Id,
		Check_In:         rs.Check_In,
		Check_Out:        rs.Check_Out,
		CreateAt:         rs.CreateAt,
		UpdateAt:         rs.UpdateAt,
	}
}

func FromDomain(domain reservation.Domain) Reservation {
	return Reservation{
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

func AllReservation(datars []Reservation) []reservation.Domain {
	All := []reservation.Domain{}
	for _, v := range datars {
		All = append(All, v.ToDomain())
	}
	return All
}
