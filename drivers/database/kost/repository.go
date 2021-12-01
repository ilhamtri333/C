package kost

import (
	"cleanarch/business/kost"
	"context"
	"errors"

	"gorm.io/gorm"
)

type KostRepository struct {
	db *gorm.DB
}

func NewKostRepository(gormDb *gorm.DB) kost.KostRepoInterface {
	return &KostRepository{
		db: gormDb,
	}
}

func (repo *KostRepository) InsertKost(ctx context.Context, domain *kost.Domain) (kost.Domain, error) {
	kostDb := FromDomain(*domain)
	err := repo.db.Create(&kostDb)
	if err.Error != nil {
		return kost.Domain{}, err.Error
	}
	return kostDb.ToDomain(), nil
}

func (repo *KostRepository) GetAllKost(ctx context.Context) ([]kost.Domain, error) {
	var kostDb []Kost
	err := repo.db.Find(&kostDb)
	if err.Error != nil {
		return []kost.Domain{}, err.Error
	}
	return AllKost(kostDb), nil
}

func (repo *KostRepository) GetKostById(ctx context.Context, id uint) (kost.Domain, error) {
	var kostDb Kost
	err := repo.db.Find(&kostDb, "id = ?", id)
	if err.Error != nil {
		return kost.Domain{}, err.Error
	}
	return kostDb.ToDomain(), nil
}

func (repo *KostRepository) UpdateKost(ctx context.Context, domain kost.Domain, id uint) (kost.Domain, error) {
	kostDb := FromDomain(domain)
	if repo.db.Save(&kostDb).Error != nil {
		return kost.Domain{}, errors.New("ga ketemu")
	}
	return kostDb.ToDomain(), nil
}

func (repo *KostRepository) DeleteKost(ctx context.Context, id uint) error {
	kostDb := Kost{}
	err := repo.db.Delete(&kostDb, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("data ga ada")
	}
	return nil
}
