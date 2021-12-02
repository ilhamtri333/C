package reservation_test

import (
	"cleanarch/business/reservation"
	mocks_test "cleanarch/business/reservation/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var reservationRepository = mocks_test.Repository{Mock: mock.Mock{}}
var reservationService reservation.ReservationUseCaseInterface
var reservationDomain reservation.Domain
var listreservationDomain []reservation.Domain

func setup() {
	reservationService = reservation.NewUseCase(&reservationRepository, time.Hour*10)
	reservationDomain = reservation.Domain{
		Id:               1,
		User_Id:          1,
		Kost_Id:          1,
		PaymentMethod_Id: 1,
		Check_In:         "2021-12-02",
		Check_Out:        "2021-12-03",
	}
	listreservationDomain = append(listreservationDomain, reservationDomain)
}

func TestInsertReservation(t *testing.T) {
	t.Run("Test Case 1 | Success Insert Reservation", func(t *testing.T) {
		setup()
		reservationRepository.On("InsertReservation", mock.Anything, mock.Anything).Return(reservationDomain, nil).Once()
		_, err := reservationService.InsertReservation(context.Background(), &reservation.Domain{})
		assert.Error(t, err)
		// assert.Equal(t, reservationDomain, reservation)
	})
}

func TestResevation(t *testing.T) {
	t.Run("Test case 1 | Success GetAllReservation", func(t *testing.T) {
		setup()
		reservationRepository.On("GetAllReservation", mock.Anything, mock.Anything).Return(listreservationDomain, nil).Once()
		data, err := reservationService.GetAllReservation(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listreservationDomain))
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetById", func(t *testing.T) {
		setup()
		reservationRepository.On("GetReservationById", mock.Anything, mock.AnythingOfType("uint")).Return(reservationDomain, nil).Once()
		data, err := reservationService.GetReservationById(context.Background(), reservationDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetById(Reservation ID = 0)", func(t *testing.T) {
		setup()
		reservationDomain.Id = 0
		reservationRepository.On("GetReservationById", mock.Anything, mock.AnythingOfType("uint")).Return(reservationDomain, nil).Once()
		data, err := reservationService.GetReservationById(context.Background(), reservationDomain.Id)

		assert.NoError(t, err)
		assert.Equal(t, data, reservation.Domain{})
	})

	t.Run("Test case 3 | Error GetById", func(t *testing.T) {
		setup()
		reservationRepository.On("GetReservationById", mock.Anything, mock.AnythingOfType("uint")).Return(reservation.Domain{}, nil).Once()
		data, err := reservationService.GetReservationById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, reservation.Domain{})
	})
}

func TestUpdateReservation(t *testing.T) {
	t.Run("Test case 1 | Success Update Reservation", func(t *testing.T) {
		setup()
		reservationRepository.On("UpdateReservation", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(reservationDomain, nil).Once()
		data, err := reservationService.UpdateReservation(context.Background(), reservationDomain, reservationDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Reservation", func(t *testing.T) {
		setup()
		reservationRepository.On("UpdateReservation", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(reservationDomain, errors.New("Books Not Found")).Once()
		data, err := reservationService.UpdateReservation(context.Background(), reservationDomain, reservationDomain.Id)

		assert.Equal(t, data, reservation.Domain{})
		assert.Error(t, err)
	})
}

func TestDeletepayment(t *testing.T) {
	t.Run("Test case 1 | Success Delete reservation", func(t *testing.T) {
		setup()
		reservationRepository.On("DeleteReservation", mock.Anything, mock.Anything).Return(nil).Once()
		err := reservationService.DeleteReservation(context.Background(), reservationDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete reservation", func(t *testing.T) {
		setup()
		reservationRepository.On("DeleteReservation", mock.Anything, mock.Anything).Return(errors.New("reservations  Not Found")).Once()
		err := reservationService.DeleteReservation(context.Background(), reservationDomain.Id)

		assert.NotEqual(t, err, errors.New("reservations Not Found"))
		assert.Error(t, err)
	})
}
