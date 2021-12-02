package detailfas_test

import (
	detailfas "cleanarch/business/detail_fas"
	mocks_test "cleanarch/business/detail_fas/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var DetailFasRepository = mocks_test.Repository{Mock: mock.Mock{}}
var DetailFasService detailfas.DetailFasUseCaseInterface
var DetailFasDomain detailfas.Domain
var listDetailFasDomain []detailfas.Domain

func setup() {
	DetailFasService = detailfas.NewUseCase(&DetailFasRepository, time.Hour*10)
	DetailFasDomain = detailfas.Domain{
		ID:        1,
		KostID:    1,
		KostFasID: 1,
	}
	listDetailFasDomain = append(listDetailFasDomain, DetailFasDomain)
}

func TestInsertDetailFas(t *testing.T) {
	t.Run("Test Case 1 | Success Insert DetailFas", func(t *testing.T) {
		setup()
		DetailFasRepository.On("InsertKost", mock.Anything, mock.Anything).Return(DetailFasDomain, nil).Once()
		_, err := DetailFasService.InsertDetailFas(context.Background(), &detailfas.Domain{})
		assert.Error(t, err)
		// assert.Equal(t, kostDomain, kost)
	})
}

func TestDetailFas(t *testing.T) {
	t.Run("Test case 1 | Success GetAllDetailFas", func(t *testing.T) {
		setup()
		DetailFasRepository.On("GetAllDetailFas", mock.Anything, mock.Anything).Return(listDetailFasDomain, nil).Once()
		data, err := DetailFasService.GetAllDetailFas(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listDetailFasDomain))
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetById", func(t *testing.T) {
		setup()
		DetailFasRepository.On("GetDetailFasById", mock.Anything, mock.AnythingOfType("uint")).Return(DetailFasDomain, nil).Once()
		data, err := DetailFasService.GetDetailFasById(context.Background(), DetailFasDomain.ID)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetById(kost ID = 0)", func(t *testing.T) {
		setup()
		DetailFasDomain.ID = 0
		DetailFasRepository.On("GetDetailFasById", mock.Anything, mock.AnythingOfType("uint")).Return(DetailFasDomain, nil).Once()
		data, err := DetailFasService.GetDetailFasById(context.Background(), DetailFasDomain.ID)

		assert.NoError(t, err)
		assert.Equal(t, data, detailfas.Domain{})
	})

	t.Run("Test case 3 | Error ById", func(t *testing.T) {
		setup()
		DetailFasRepository.On("GetDetailFasById", mock.Anything, mock.AnythingOfType("uint")).Return(detailfas.Domain{}, nil).Once()
		data, err := DetailFasService.GetDetailFasById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, detailfas.Domain{})
	})
}

func TestUpdateDetailFas(t *testing.T) {
	t.Run("Test case 1 | Success Update DetailFas", func(t *testing.T) {
		setup()
		DetailFasRepository.On("UpdateDetailFas", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(DetailFasDomain, nil).Once()
		data, err := DetailFasService.UpdateDetailFas(context.Background(), DetailFasDomain, DetailFasDomain.ID)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update DetailFas", func(t *testing.T) {
		setup()
		DetailFasRepository.On("UpdateDetailFas", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(DetailFasDomain, errors.New("Books Not Found")).Once()
		data, err := DetailFasService.UpdateDetailFas(context.Background(), DetailFasDomain, DetailFasDomain.ID)

		assert.Equal(t, data, detailfas.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteDetailFas(t *testing.T) {
	t.Run("Test case 1 | Success Delete DetailFas", func(t *testing.T) {
		setup()
		DetailFasRepository.On("DeleteDetailFas", mock.Anything, mock.Anything).Return(nil).Once()
		err := DetailFasService.DeleteDetailFas(context.Background(), DetailFasDomain.ID)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete DetailFas", func(t *testing.T) {
		setup()
		DetailFasRepository.On("DeleteDetailFas", mock.Anything, mock.Anything).Return(errors.New("DetailFass  Not Found")).Once()
		err := DetailFasService.DeleteDetailFas(context.Background(), DetailFasDomain.ID)

		assert.NotEqual(t, err, errors.New("DetailFas Not Found"))
		assert.Error(t, err)
	})
}
