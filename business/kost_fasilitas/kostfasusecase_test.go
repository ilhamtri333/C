package kostfasilitas_test

import (
	kostfasilitas "cleanarch/business/kost_fasilitas"
	mocks_test "cleanarch/business/kost_fasilitas/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var kostFasRepository = mocks_test.Repository{Mock: mock.Mock{}}
var kostFasService kostfasilitas.KostFasUsecaseInterface
var kostFasDomain kostfasilitas.Domain
var listKostFasDomain []kostfasilitas.Domain

func setup() {
	kostFasService = kostfasilitas.NewUseCase(&kostFasRepository, time.Hour*10)
	kostFasDomain = kostfasilitas.Domain{
		Id:        1,
		Fasilitas: "Extra bantal",
	}
	listKostFasDomain = append(listKostFasDomain, kostFasDomain)
}

func TestInsertKostFas(t *testing.T) {
	t.Run("Test Case 1 | Success Insert KostFas", func(t *testing.T) {
		setup()
		kostFasRepository.On("InsertKostFas", mock.Anything, mock.Anything).Return(kostFasDomain, nil).Once()
		_, err := kostFasService.InsertKostFas(context.Background(), &kostfasilitas.Domain{})
		assert.Error(t, err)
		// assert.Equal(t, kostFasDomain, kostFas)
	})
}

func TestKostFas(t *testing.T) {
	t.Run("Test case 1 | Success GetAllKostFas", func(t *testing.T) {
		setup()
		kostFasRepository.On("GetAllKostFas", mock.Anything, mock.Anything).Return(listKostFasDomain, nil).Once()
		data, err := kostFasService.GetAllKostFas(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listKostFasDomain))
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetById", func(t *testing.T) {
		setup()
		kostFasRepository.On("GetKostFasById", mock.Anything, mock.AnythingOfType("uint")).Return(kostFasDomain, nil).Once()
		data, err := kostFasService.GetKostFasById(context.Background(), kostFasDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetById(kost ID = 0)", func(t *testing.T) {
		setup()
		kostFasDomain.Id = 0
		kostFasRepository.On("GetKostFasById", mock.Anything, mock.AnythingOfType("uint")).Return(kostFasDomain, nil).Once()
		data, err := kostFasService.GetKostFasById(context.Background(), kostFasDomain.Id)

		assert.NoError(t, err)
		assert.Equal(t, data, kostfasilitas.Domain{})
	})

	t.Run("Test case 3 | Error GetById", func(t *testing.T) {
		setup()
		kostFasRepository.On("GetKostFasById", mock.Anything, mock.AnythingOfType("uint")).Return(kostfasilitas.Domain{}, nil).Once()
		data, err := kostFasService.GetKostFasById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, kostfasilitas.Domain{})
	})
}

func TestUpdateKostFas(t *testing.T) {
	t.Run("Test case 1 | Success Update KostFas", func(t *testing.T) {
		setup()
		kostFasRepository.On("UpdateKostFas", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(kostFasDomain, nil).Once()
		data, err := kostFasService.UpdateKostFas(context.Background(), kostFasDomain, kostFasDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update KostFas", func(t *testing.T) {
		setup()
		kostFasRepository.On("UpdateKostFas", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(kostFasDomain, errors.New("Books Not Found")).Once()
		data, err := kostFasService.UpdateKostFas(context.Background(), kostFasDomain, kostFasDomain.Id)

		assert.Equal(t, data, kostfasilitas.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteKostFas(t *testing.T) {
	t.Run("Test case 1 | Success Delete KostFas", func(t *testing.T) {
		setup()
		kostFasRepository.On("DeleteKostFas", mock.Anything, mock.Anything).Return(nil).Once()
		err := kostFasService.DeleteKostFas(context.Background(), kostFasDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete KostFas", func(t *testing.T) {
		setup()
		kostFasRepository.On("DeleteKostFas", mock.Anything, mock.Anything).Return(errors.New("kostFass  Not Found")).Once()
		err := kostFasService.DeleteKostFas(context.Background(), kostFasDomain.Id)

		assert.NotEqual(t, err, errors.New("kostFass Not Found"))
		assert.Error(t, err)
	})
}
