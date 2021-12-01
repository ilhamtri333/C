package kost

import (
	"context"
	"errors"
	"time"
)

type KostUsecase struct {
	repo KostRepoInterface
	ctx  time.Duration
}

func NewUseCase(repoKost KostRepoInterface, contextTimeout time.Duration) KostUseCaseInterface {
	return &KostUsecase{
		repo: repoKost,
		ctx:  contextTimeout,
	}
}

func (usecase *KostUsecase) InsertKost(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Nama == "" {
		return Domain{}, errors.New("nama empty")
	}
	if domain.Type_Kost == "" {
		return Domain{}, errors.New("type kost empty")
	}
	if domain.Desc_Kost == "" {
		return Domain{}, errors.New("desc kost empty")
	}
	if domain.Address == "" {
		return Domain{}, errors.New("address empty")
	}
	if domain.Duration == 0 {
		return Domain{}, errors.New("duration empty")
	}
	if domain.OwnerID == 0 {
		return Domain{}, errors.New("owner id empty")
	}
	kost, err := usecase.repo.InsertKost(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return kost, nil
}

func (usecase *KostUsecase) GetAllKost(ctx context.Context) ([]Domain, error) {
	return usecase.repo.GetAllKost(ctx)
}

func (usecase *KostUsecase) GetKostById(ctx context.Context, id uint) (Domain, error) {
	kost, err := usecase.repo.GetKostById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if kost.ID == 0 {
		return Domain{}, err
	}
	return kost, nil
}

func (usecase *KostUsecase) UpdateKost(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.ID = (id)
	kost, err := usecase.repo.UpdateKost(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return kost, nil
}

func (usecase *KostUsecase) DeleteKost(ctx context.Context, id uint) error {
	err := usecase.repo.DeleteKost(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
