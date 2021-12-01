package payment

import (
	"context"
	"errors"
	"time"
)

type Payment_MethodUseCase struct {
	repo PaymentMetodRepoInterface
	ctx  time.Duration
}

func NewUseCase(repoPayment PaymentMetodRepoInterface, contextTimeout time.Duration) PaymentMetodUsecaseInterface {
	return &Payment_MethodUseCase{
		repo: repoPayment,
		ctx:  contextTimeout,
	}
}

func (usecase *Payment_MethodUseCase) InsertPayment(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Type == "" {
		return Domain{}, errors.New("method payment empty")
	}
	if domain.Nama == "" {
		return Domain{}, errors.New("nama method payment empty")
	}

	payment_method, err := usecase.repo.InsertPayment(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return payment_method, nil
}

func (usecase *Payment_MethodUseCase) GetAllPayment(ctx context.Context) ([]Domain, error) {
	return usecase.repo.GetAllPayment(ctx)
}

func (usecase *Payment_MethodUseCase) GetPaymentById(ctx context.Context, id uint) (Domain, error) {
	payment_method, err := usecase.repo.GetPaymentById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if payment_method.Id == 0 {
		return Domain{}, err
	}
	return payment_method, nil
}

func (usecase *Payment_MethodUseCase) UpdatePayment(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	payment_method, err := usecase.repo.UpdatePayment(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return payment_method, nil
}

func (usecase *Payment_MethodUseCase) DeletePayment(ctx context.Context, id uint) error {
	err := usecase.repo.DeletePayment(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
