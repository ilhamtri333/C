package detailfas

import (
	detailfas "cleanarch/business/detail_fas"
	"context"
	"errors"

	"gorm.io/gorm"
)

type DetailFasRepository struct {
	db *gorm.DB
}

func NewDetailFasRepository(gormDb *gorm.DB) detailfas.DetailFasRepoInterface {
	return &DetailFasRepository{
		db: gormDb,
	}
}

func (repo *DetailFasRepository) InsertDetailFas(ctx context.Context, domain *detailfas.Domain) (detailfas.Domain, error) {
	dfDb := FromDomain(*domain)
	err := repo.db.Create(&dfDb)
	if err.Error != nil {
		return detailfas.Domain{}, err.Error
	}
	return dfDb.ToDomain(), nil
}

func (repo *DetailFasRepository) GetAllDetailFas(ctx context.Context) ([]detailfas.Domain, error) {
	var dfDb []DetailFas
	err := repo.db.Find(&dfDb)
	if err.Error != nil {
		return []detailfas.Domain{}, err.Error
	}
	return AllDetailFas(dfDb), nil
}

func (repo *DetailFasRepository) GetDetailFasById(ctx context.Context, id uint) (detailfas.Domain, error) {
	var dfDb DetailFas
	err := repo.db.Find(&dfDb, "id = ?", id)
	if err.Error != nil {
		return detailfas.Domain{}, err.Error
	}
	return dfDb.ToDomain(), nil
}

func (repo *DetailFasRepository) UpdateDetailFas(ctx context.Context, domain detailfas.Domain, id uint) (detailfas.Domain, error) {
	dfDb := FromDomain(domain)
	if repo.db.Save(&dfDb).Error != nil {
		return detailfas.Domain{}, errors.New("ga ketemu")
	}
	return dfDb.ToDomain(), nil
}

func (repo *DetailFasRepository) DeleteDetailFas(ctx context.Context, id uint) error {
	dfDb := DetailFas{}
	err := repo.db.Delete(&dfDb, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("data ga ada")
	}
	return nil
}
