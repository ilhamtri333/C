package kostfasilitas

import (
	kostfasilitas "cleanarch/business/kost_fasilitas"
	"context"
	"errors"

	"gorm.io/gorm"
)

type KostFasRepository struct {
	db *gorm.DB
}

func NewKostFasRepository(gormDb *gorm.DB) kostfasilitas.KostFasMetodRepoInterface {
	return &KostFasRepository{
		db: gormDb,
	}
}

func (repo *KostFasRepository) InsertKostFas(ctx context.Context, domain *kostfasilitas.Domain) (kostfasilitas.Domain, error) {
	fasDb := FromDomain(*domain)
	err := repo.db.Create(&fasDb)
	if err.Error != nil {
		return kostfasilitas.Domain{}, err.Error
	}
	return fasDb.ToDomain(), nil
}

func (repo *KostFasRepository) GetAllKostFas(ctx context.Context) ([]kostfasilitas.Domain, error) {
	var fasDb []KostFas
	err := repo.db.Find(&fasDb)
	if err.Error != nil {
		return []kostfasilitas.Domain{}, err.Error
	}
	return AllKostFas(fasDb), nil
}

func (repo *KostFasRepository) GetKostFasById(ctx context.Context, id uint) (kostfasilitas.Domain, error) {
	var fasDb KostFas
	err := repo.db.Find(&fasDb, "id = ?", id)
	if err.Error != nil {
		return kostfasilitas.Domain{}, err.Error
	}
	return fasDb.ToDomain(), nil
}

func (repo *KostFasRepository) UpdateKostFas(ctx context.Context, domain kostfasilitas.Domain, id uint) (kostfasilitas.Domain, error) {
	fasDb := FromDomain(domain)
	if repo.db.Save(&fasDb).Error != nil {
		return kostfasilitas.Domain{}, errors.New("ga ketemu")
	}
	return fasDb.ToDomain(), nil
}

func (repo *KostFasRepository) DeleteKostFas(ctx context.Context, id uint) error {
	fasDb := KostFas{}
	err := repo.db.Delete(&fasDb, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("data ga ada")
	}
	return nil
}
