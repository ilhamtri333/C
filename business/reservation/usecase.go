package reservation

import (
	timing "cleanarch/helper/time"
	"context"
	"errors"
	"time"
)

type ReservationUseCase struct {
	repo ReservationRepoInterface
	ctx  time.Duration
}

func NewUseCase(repoReserv ReservationRepoInterface, contextTimeout time.Duration) ReservationUseCaseInterface {
	return &ReservationUseCase{
		repo: repoReserv,
		ctx:  contextTimeout,
	}
}

func (usecase *ReservationUseCase) InsertReservation(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.User_Id == 0 {
		return Domain{}, errors.New("userid empty")
	}
	if domain.Kost_Id == 0 {
		return Domain{}, errors.New("kostid empty")
	}
	if domain.PaymentMethod_Id == 0 {
		return Domain{}, errors.New("payment empty")
	}
	if !timing.TimingCheck(domain.Check_In) {
		return Domain{}, errors.New("checkin empty")
	}
	if !timing.TimingCheck(domain.Check_Out) {
		return Domain{}, errors.New("checkout empty")
	}

	reserv, err := usecase.repo.InsertReservation(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return reserv, nil
}

func (usecase *ReservationUseCase) GetAllReservation(ctx context.Context) ([]Domain, error) {
	return usecase.repo.GetAllReservation(ctx)
}

func (usecase *ReservationUseCase) GetReservationById(ctx context.Context, id uint) (Domain, error) {
	reserv, err := usecase.repo.GetReservationById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if reserv.Id == 0 {
		return Domain{}, err
	}
	return reserv, nil
}

func (usecase *ReservationUseCase) UpdateReservation(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	reserv, err := usecase.repo.UpdateReservation(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return reserv, nil
}

func (usecase *ReservationUseCase) DeleteReservation(ctx context.Context, id uint) error {
	err := usecase.repo.DeleteReservation(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
