package users

import (
	"cleanarch/business/users"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db: gormDb,
	}
}
func (repo *UserRepository) Register(domain *users.Domain, ctx context.Context) (users.Domain, error) {
	userDb := FromDomain(*domain)
	err := repo.db.Create(&userDb)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return userDb.ToDomain(), nil

}

func (repo *UserRepository) Login(domain users.Domain, ctx context.Context) (users.Domain, error) {
	userDb := FromDomain(domain)

	err := repo.db.Where("email = ? AND password = ?", userDb.Email, userDb.Password).First(&userDb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users.Domain{}, errors.New("email not found")
		}
		return users.Domain{}, errors.New("error in database")
	}
	return userDb.ToDomain(), nil
}

func (repo *UserRepository) GetAllUser(ctx context.Context) ([]users.Domain, error) {
	var userDb []User
	err := repo.db.Find(&userDb)
	if err.Error != nil {
		return []users.Domain{}, err.Error
	}
	return ListUserToDomain(userDb), nil
}

func (repo *UserRepository) GetUserById(ctx context.Context, id uint) (users.Domain, error) {
	var userDb User
	err := repo.db.Find(&userDb, "id = ?", id)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return userDb.ToDomain(), nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, domain users.Domain, id uint) (users.Domain, error) {
	userDb := FromDomain(domain)
	if repo.db.Save(&userDb).Error != nil {
		return users.Domain{}, errors.New("ga ketemu")
	}
	return userDb.ToDomain(), nil
}

func (repo *UserRepository) DeleteUser(ctx context.Context, id uint) error {
	userDb := User{}
	err := repo.db.Delete(&userDb, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("data ga ada")
	}
	return nil
}
