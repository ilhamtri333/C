package reservation

import (
	"cleanarch/business/reservation"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(gormDb *gorm.DB) reservation.ReservationRepoInterface {
	return &ReservationRepository{
		db: gormDb,
	}
}

func (repo *ReservationRepository) InsertReservation(ctx context.Context, domain *reservation.Domain) (reservation.Domain, error) {
	reservDb := FromDomain(*domain)
	err := repo.db.Create(&reservDb)
	if err.Error != nil {
		return reservation.Domain{}, err.Error
	}
	return reservDb.ToDomain(), nil
}

func (repo *ReservationRepository) GetAllReservation(ctx context.Context) ([]reservation.Domain, error) {
	var reservDb []Reservation
	err := repo.db.Find(&reservDb)
	if err.Error != nil {
		return []reservation.Domain{}, err.Error
	}
	return AllReservation(reservDb), nil
}

func (repo *ReservationRepository) GetReservationById(ctx context.Context, id uint) (reservation.Domain, error) {
	var reservDb Reservation
	err := repo.db.Find(&reservDb, "id = ?", id)
	if err.Error != nil {
		return reservation.Domain{}, err.Error
	}
	return reservDb.ToDomain(), nil
}

func (repo *ReservationRepository) UpdateReservation(ctx context.Context, domain reservation.Domain, id uint) (reservation.Domain, error) {
	reservDb := FromDomain(domain)
	if repo.db.Save(&reservDb).Error != nil {
		return reservation.Domain{}, errors.New("ga ketemu")
	}
	return reservDb.ToDomain(), nil
}

func (repo *ReservationRepository) DeleteReservation(ctx context.Context, id uint) error {
	reservDb := Reservation{}
	err := repo.db.Delete(&reservDb, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("data ga ada")
	}
	return nil
}
