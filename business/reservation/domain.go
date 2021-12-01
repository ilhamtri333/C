package reservation

import (
	"context"
	"time"
)

type Domain struct {
	Id      uint
	User_Id uint
	// User             users.Domain
	Kost_Id uint
	// Kost             kost.Domain
	PaymentMethod_Id uint
	// PaymentMethod    paymentmethod.Domain
	Check_In  string
	Check_Out string
	CreateAt  time.Time
	UpdateAt  time.Time
}

type ReservationUseCaseInterface interface {
	InsertReservation(ctx context.Context, domain *Domain) (Domain, error)
	GetAllReservation(ctx context.Context) ([]Domain, error)
	GetReservationById(ctx context.Context, id uint) (Domain, error)
	UpdateReservation(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteReservation(ctx context.Context, id uint) error
}

type ReservationRepoInterface interface {
	InsertReservation(ctx context.Context, domain *Domain) (Domain, error)
	GetAllReservation(ctx context.Context) ([]Domain, error)
	GetReservationById(ctx context.Context, id uint) (Domain, error)
	UpdateReservation(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeleteReservation(ctx context.Context, id uint) error
}
