package kost_test

import (
	"cleanarch/business/kost"
	mocks_test "cleanarch/business/kost/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var kostRepository = mocks_test.Repository{Mock: mock.Mock{}}
var kostService kost.KostUseCaseInterface
var kostDomain kost.Domain
var listKostDomain []kost.Domain

func setup() {
	kostService = kost.NewUseCase(&kostRepository, time.Hour*10)
	kostDomain = kost.Domain{
		ID:        1,
		Nama:      "ilham",
		Type_Kost: "kamar",
		Desc_Kost: "ada tempat tidur",
		Address:   "jl slamet",
		Duration:  1,
		OwnerID:   1,
	}
	listKostDomain = append(listKostDomain, kostDomain)
}

func TestInsertKost(t *testing.T) {
	t.Run("Test Case 1 | Success Insert Kost", func(t *testing.T) {
		setup()
		kostRepository.On("InsertKost", mock.Anything, mock.Anything).Return(kostDomain, nil).Once()
		_, err := kostService.InsertKost(context.Background(), &kost.Domain{})
		assert.Error(t, err)
		// assert.Equal(t, kostDomain, kost)
	})
}

func TestKost(t *testing.T) {
	t.Run("Test case 1 | Success GetAllKost", func(t *testing.T) {
		setup()
		kostRepository.On("GetAllKost", mock.Anything, mock.Anything).Return(listKostDomain, nil).Once()
		data, err := kostService.GetAllKost(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listKostDomain))
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetById", func(t *testing.T) {
		setup()
		kostRepository.On("GetKostById", mock.Anything, mock.AnythingOfType("uint")).Return(kostDomain, nil).Once()
		data, err := kostService.GetKostById(context.Background(), kostDomain.ID)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetById(kost ID = 0)", func(t *testing.T) {
		setup()
		kostDomain.ID = 0
		kostRepository.On("GetKostById", mock.Anything, mock.AnythingOfType("uint")).Return(kostDomain, nil).Once()
		data, err := kostService.GetKostById(context.Background(), kostDomain.ID)

		assert.NoError(t, err)
		assert.Equal(t, data, kost.Domain{})
	})

	t.Run("Test case 3 | Error SearchBookById", func(t *testing.T) {
		setup()
		kostRepository.On("GetKostById", mock.Anything, mock.AnythingOfType("uint")).Return(kost.Domain{}, nil).Once()
		data, err := kostService.GetKostById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, kost.Domain{})
	})
}

func TestUpdateBook(t *testing.T) {
	t.Run("Test case 1 | Success Update Kost", func(t *testing.T) {
		setup()
		kostRepository.On("UpdateKost", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(kostDomain, nil).Once()
		data, err := kostService.UpdateKost(context.Background(), kostDomain, kostDomain.ID)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Kost", func(t *testing.T) {
		setup()
		kostRepository.On("UpdateKost", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(kostDomain, errors.New("Books Not Found")).Once()
		data, err := kostService.UpdateKost(context.Background(), kostDomain, kostDomain.ID)

		assert.Equal(t, data, kost.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteKost(t *testing.T) {
	t.Run("Test case 1 | Success Delete Kost", func(t *testing.T) {
		setup()
		kostRepository.On("DeleteKost", mock.Anything, mock.Anything).Return(nil).Once()
		err := kostService.DeleteKost(context.Background(), kostDomain.ID)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete DetailFas", func(t *testing.T) {
		setup()
		kostRepository.On("DeleteKost", mock.Anything, mock.Anything).Return(errors.New("kosts  Not Found")).Once()
		err := kostService.DeleteKost(context.Background(), kostDomain.ID)

		assert.NotEqual(t, err, errors.New("kosts Not Found"))
		assert.Error(t, err)
	})
}
