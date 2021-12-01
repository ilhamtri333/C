package payment

import (
	payment "cleanarch/business/payment_method"
	"context"
	"errors"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(gormDb *gorm.DB) payment.PaymentMetodRepoInterface {
	return &PaymentRepository{
		db: gormDb,
	}
}

func (repo *PaymentRepository) InsertPayment(ctx context.Context, domain *payment.Domain) (payment.Domain, error) {
	payDb := FromDomain(*domain)
	err := repo.db.Create(&payDb)
	if err.Error != nil {
		return payment.Domain{}, err.Error
	}
	return payDb.ToDomain(), nil
}

func (repo *PaymentRepository) GetAllPayment(ctx context.Context) ([]payment.Domain, error) {
	var payDb []Payment
	err := repo.db.Find(&payDb)
	if err.Error != nil {
		return []payment.Domain{}, err.Error
	}
	return AllPayment(payDb), nil
}

func (repo *PaymentRepository) GetPaymentById(ctx context.Context, id uint) (payment.Domain, error) {
	var payDb Payment
	err := repo.db.Find(&payDb, "id = ?", id)
	if err.Error != nil {
		return payment.Domain{}, err.Error
	}
	return payDb.ToDomain(), nil
}

func (repo *PaymentRepository) UpdatePayment(ctx context.Context, domain payment.Domain, id uint) (payment.Domain, error) {
	payDb := FromDomain(domain)
	if repo.db.Save(&payDb).Error != nil {
		return payment.Domain{}, errors.New("ga ketemu")
	}
	return payDb.ToDomain(), nil
}

func (repo *PaymentRepository) DeletePayment(ctx context.Context, id uint) error {
	payDb := Payment{}
	err := repo.db.Delete(&payDb, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("data ga ada")
	}
	return nil
}
