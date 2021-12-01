package payment

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	Type      string
	Nama      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentMetodUsecaseInterface interface {
	InsertPayment(ctx context.Context, domain *Domain) (Domain, error)
	GetAllPayment(ctx context.Context) ([]Domain, error)
	GetPaymentById(ctx context.Context, id uint) (Domain, error)
	UpdatePayment(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeletePayment(ctx context.Context, id uint) error
}

type PaymentMetodRepoInterface interface {
	InsertPayment(ctx context.Context, domain *Domain) (Domain, error)
	GetAllPayment(ctx context.Context) ([]Domain, error)
	GetPaymentById(ctx context.Context, id uint) (Domain, error)
	UpdatePayment(ctx context.Context, domain Domain, id uint) (Domain, error)
	DeletePayment(ctx context.Context, id uint) error
}
