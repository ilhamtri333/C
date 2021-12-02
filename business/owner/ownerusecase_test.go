package owner_test

import (
	"cleanarch/business/owner"
	mocks_test "cleanarch/business/owner/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ownerRepository = mocks_test.Repository{Mock: mock.Mock{}}
var ownerService owner.OwnerUseCaseInterface
var ownerDomain owner.Domain
var listownerDomain []owner.Domain

func setup() {
	ownerService = owner.NewUseCase(&ownerRepository, time.Hour*10)
	ownerDomain = owner.Domain{
		ID:      1,
		Owner:   "ilham",
		Email:   "ilham@gmail.com",
		No_telp: 821821821821,
	}
	listownerDomain = append(listownerDomain, ownerDomain)
}

func TestInsertOwner(t *testing.T) {
	t.Run("Test Case 1 | Success Insert Owner", func(t *testing.T) {
		setup()
		ownerRepository.On("InsertOwner", mock.Anything, mock.Anything).Return(ownerDomain, nil).Once()
		_, err := ownerService.InsertOwner(context.Background(), &owner.Domain{})
		assert.Error(t, err)
		// assert.Equal(t, ownerDomain, owner)
	})
}

func TestOwner(t *testing.T) {
	t.Run("Test case 1 | Success GetAllOwner", func(t *testing.T) {
		setup()
		ownerRepository.On("GetAllOwner", mock.Anything, mock.Anything).Return(listownerDomain, nil).Once()
		data, err := ownerService.GetAllOwner(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listownerDomain))
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetById", func(t *testing.T) {
		setup()
		ownerRepository.On("GetOwnerById", mock.Anything, mock.AnythingOfType("uint")).Return(ownerDomain, nil).Once()
		data, err := ownerService.GetOwnerById(context.Background(), ownerDomain.ID)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetById(owner ID = 0)", func(t *testing.T) {
		setup()
		ownerDomain.ID = 0
		ownerRepository.On("GetOwnerById", mock.Anything, mock.AnythingOfType("uint")).Return(ownerDomain, nil).Once()
		data, err := ownerService.GetOwnerById(context.Background(), ownerDomain.ID)

		assert.NoError(t, err)
		assert.Equal(t, data, owner.Domain{})
	})

	t.Run("Test case 3 | Error SearchBookById", func(t *testing.T) {
		setup()
		ownerRepository.On("GetOwnerById", mock.Anything, mock.AnythingOfType("uint")).Return(owner.Domain{}, nil).Once()
		data, err := ownerService.GetOwnerById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, owner.Domain{})
	})
}

func TestUpdateBook(t *testing.T) {
	t.Run("Test case 1 | Success Update Owner", func(t *testing.T) {
		setup()
		ownerRepository.On("UpdateOwner", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(ownerDomain, nil).Once()
		data, err := ownerService.UpdateOwner(context.Background(), ownerDomain, ownerDomain.ID)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Owner", func(t *testing.T) {
		setup()
		ownerRepository.On("UpdateOwner", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(ownerDomain, errors.New("Books Not Found")).Once()
		data, err := ownerService.UpdateOwner(context.Background(), ownerDomain, ownerDomain.ID)

		assert.Equal(t, data, owner.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteOwner(t *testing.T) {
	t.Run("Test case 1 | Success Delete Owner", func(t *testing.T) {
		setup()
		ownerRepository.On("DeleteOwner", mock.Anything, mock.Anything).Return(nil).Once()
		err := ownerService.DeleteOwner(context.Background(), ownerDomain.ID)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete Owner", func(t *testing.T) {
		setup()
		ownerRepository.On("DeleteOwner", mock.Anything, mock.Anything).Return(errors.New("owners  Not Found")).Once()
		err := ownerService.DeleteOwner(context.Background(), ownerDomain.ID)

		assert.NotEqual(t, err, errors.New("owners Not Found"))
		assert.Error(t, err)
	})
}
