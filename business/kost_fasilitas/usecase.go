package kostfasilitas

import (
	"context"
	"errors"
	"time"
)

type KostFasUseCase struct {
	repo KostFasMetodRepoInterface
	ctx  time.Duration
}

func NewUseCase(repoKostFas KostFasMetodRepoInterface, contextTimeout time.Duration) KostFasUsecaseInterface {
	return &KostFasUseCase{
		repo: repoKostFas,
		ctx:  contextTimeout,
	}
}

func (usecase *KostFasUseCase) InsertKostFas(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Fasilitas == "" {
		return Domain{}, errors.New("fasilitas empty")
	}

	kostfas, err := usecase.repo.InsertKostFas(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return kostfas, nil
}

func (usecase *KostFasUseCase) GetAllKostFas(ctx context.Context) ([]Domain, error) {
	return usecase.repo.GetAllKostFas(ctx)
}

func (usecase *KostFasUseCase) GetKostFasById(ctx context.Context, id uint) (Domain, error) {
	kostfas, err := usecase.repo.GetKostFasById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if kostfas.Id == 0 {
		return Domain{}, err
	}
	return kostfas, nil
}

func (usecase *KostFasUseCase) UpdateKostFas(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	kostfas, err := usecase.repo.UpdateKostFas(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return kostfas, nil
}

func (usecase *KostFasUseCase) DeleteKostFas(ctx context.Context, id uint) error {
	err := usecase.repo.DeleteKostFas(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
