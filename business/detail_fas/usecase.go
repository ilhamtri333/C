package detailfas

import (
	"context"
	"errors"
	"time"
)

type DetailFasUsecase struct {
	repo DetailFasRepoInterface
	ctx  time.Duration
}

func NewUseCase(repoDetailFas DetailFasRepoInterface, contextTimeout time.Duration) DetailFasUseCaseInterface {
	return &DetailFasUsecase{
		repo: repoDetailFas,
		ctx:  contextTimeout,
	}
}

func (usecase *DetailFasUsecase) InsertDetailFas(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.KostID == 0 {
		return Domain{}, errors.New("kost id empty")
	}
	if domain.KostFasID == 0 {
		return Domain{}, errors.New("type kost empty")
	}
	kost, err := usecase.repo.InsertDetailFas(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return kost, nil
}

func (usecase *DetailFasUsecase) GetAllDetailFas(ctx context.Context) ([]Domain, error) {
	return usecase.repo.GetAllDetailFas(ctx)
}

func (usecase *DetailFasUsecase) GetDetailFasById(ctx context.Context, id uint) (Domain, error) {
	kost, err := usecase.repo.GetDetailFasById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if kost.ID == 0 {
		return Domain{}, err
	}
	return kost, nil
}

func (usecase *DetailFasUsecase) UpdateDetailFas(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.ID = (id)
	kost, err := usecase.repo.UpdateDetailFas(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return kost, nil
}

func (usecase *DetailFasUsecase) DeleteDetailFas(ctx context.Context, id uint) error {
	err := usecase.repo.DeleteDetailFas(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
