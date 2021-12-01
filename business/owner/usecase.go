package owner

import (
	"context"
	"errors"
	"time"
)

type OwnerUseCase struct {
	repo OwnerRepoInterface
	ctx  time.Duration
}

func NewUseCase(repoPayment OwnerRepoInterface, contextTimeout time.Duration) OwnerUseCaseInterface {
	return &OwnerUseCase{
		repo: repoPayment,
		ctx:  contextTimeout,
	}
}

func (usecase *OwnerUseCase) InsertOwner(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Owner == "" {
		return Domain{}, errors.New("owner empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.No_telp == 0 {
		return Domain{}, errors.New("nomor telepon empty")
	}

	owner, err := usecase.repo.InsertOwner(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return owner, nil
}

func (usecase *OwnerUseCase) GetAllOwner(ctx context.Context) ([]Domain, error) {
	return usecase.repo.GetAllOwner(ctx)
}

func (usecase *OwnerUseCase) GetOwnerById(ctx context.Context, id uint) (Domain, error) {
	owner, err := usecase.repo.GetOwnerById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if owner.ID == 0 {
		return Domain{}, err
	}
	return owner, nil
}

func (usecase *OwnerUseCase) UpdateOwner(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.ID = (id)
	owner, err := usecase.repo.UpdateOwner(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return owner, nil
}

func (usecase *OwnerUseCase) DeleteOwner(ctx context.Context, id uint) error {
	err := usecase.repo.DeleteOwner(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
