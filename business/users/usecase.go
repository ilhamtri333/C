package users

import (
	_middleware "cleanarch/app/middleware"
	"context"
	"errors"
	"fmt"
	"time"
)

type UserUseCase struct {
	//repo
	repo UserRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

func NewUseCase(userRepo UserRepoInterface, contextTimeout time.Duration, configJWT *_middleware.ConfigJWT) UserUseCaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
		jwt:  configJWT,
	}
}

func (usecase *UserUseCase) Register(domain *Domain, ctx context.Context) (Domain, error) {
	if domain.Nama == "" {
		return Domain{}, errors.New("nama empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")

	}
	user, err := usecase.repo.Register(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) Login(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	user, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	fmt.Println(user)
	user.Token = usecase.jwt.GenerateToken(user.Id)
	return user, nil
}

func (usecase *UserUseCase) GetAllUser(ctx context.Context) ([]Domain, error) {
	return usecase.repo.GetAllUser(ctx)
}

func (usecase *UserUseCase) GetUserById(ctx context.Context, id uint) (Domain, error) {
	user, err := usecase.repo.GetUserById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if user.Id == 0 {
		return Domain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) UpdateUser(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	user, err := usecase.repo.UpdateUser(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (usecase *UserUseCase) DeleteUser(ctx context.Context, id uint) error {
	err := usecase.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
