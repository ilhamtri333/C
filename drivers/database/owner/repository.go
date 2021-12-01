package owner

import (
	"cleanarch/business/owner"
	"context"
	"errors"

	"gorm.io/gorm"
)

type OwnerRepository struct {
	db *gorm.DB
}

func NewOwnerRepository(gormDb *gorm.DB) owner.OwnerRepoInterface {
	return &OwnerRepository{
		db: gormDb,
	}
}

func (repo *OwnerRepository) InsertOwner(ctx context.Context, domain *owner.Domain) (owner.Domain, error) {
	ownDb := FromDomain(*domain)
	err := repo.db.Create(&ownDb)
	if err.Error != nil {
		return owner.Domain{}, err.Error
	}
	return ownDb.ToDomain(), nil
}

func (repo *OwnerRepository) GetAllOwner(ctx context.Context) ([]owner.Domain, error) {
	var ownDb []Owner
	err := repo.db.Find(&ownDb)
	if err.Error != nil {
		return []owner.Domain{}, err.Error
	}
	return AllOwner(ownDb), nil
}

func (repo *OwnerRepository) GetOwnerById(ctx context.Context, id uint) (owner.Domain, error) {
	var ownDb Owner
	err := repo.db.Preload("Kost").Find(&ownDb, "id = ?", id)
	if err.Error != nil {
		return owner.Domain{}, err.Error
	}
	return ownDb.ToDomain(), nil
}

func (repo *OwnerRepository) UpdateOwner(ctx context.Context, domain owner.Domain, id uint) (owner.Domain, error) {
	ownDb := FromDomain(domain)
	if repo.db.Save(&ownDb).Error != nil {
		return owner.Domain{}, errors.New("ga ketemu")
	}
	return ownDb.ToDomain(), nil
}

func (repo *OwnerRepository) DeleteOwner(ctx context.Context, id uint) error {
	ownDb := Owner{}
	err := repo.db.Delete(&ownDb, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("data ga ada")
	}
	return nil
}
