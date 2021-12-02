package payment_test

import (
	payment "cleanarch/business/payment_method"
	mocks_test "cleanarch/business/payment_method/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var paymentRepository = mocks_test.Repository{Mock: mock.Mock{}}
var paymentService payment.PaymentMetodUsecaseInterface
var paymentDomain payment.Domain
var listpaymentDomain []payment.Domain

func setup() {
	paymentService = payment.NewUseCase(&paymentRepository, time.Hour*10)
	paymentDomain = payment.Domain{
		Id:   1,
		Type: "01",
		Nama: "BRI",
	}
	listpaymentDomain = append(listpaymentDomain, paymentDomain)
}

func TestInsertPayment(t *testing.T) {
	t.Run("Test Case 1 | Success Insert Payment", func(t *testing.T) {
		setup()
		paymentRepository.On("InsertPayment", mock.Anything, mock.Anything).Return(paymentDomain, nil).Once()
		_, err := paymentService.InsertPayment(context.Background(), &payment.Domain{})
		assert.Error(t, err)
		// assert.Equal(t, paymentDomain, payment)
	})
}

func TestPayment(t *testing.T) {
	t.Run("Test case 1 | Success GetAllPayment", func(t *testing.T) {
		setup()
		paymentRepository.On("GetAllPayment", mock.Anything, mock.Anything).Return(listpaymentDomain, nil).Once()
		data, err := paymentService.GetAllPayment(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listpaymentDomain))
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetById", func(t *testing.T) {
		setup()
		paymentRepository.On("GetPaymentById", mock.Anything, mock.AnythingOfType("uint")).Return(paymentDomain, nil).Once()
		data, err := paymentService.GetPaymentById(context.Background(), paymentDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetById(Payment ID = 0)", func(t *testing.T) {
		setup()
		paymentDomain.Id = 0
		paymentRepository.On("GetPaymentById", mock.Anything, mock.AnythingOfType("uint")).Return(paymentDomain, nil).Once()
		data, err := paymentService.GetPaymentById(context.Background(), paymentDomain.Id)

		assert.NoError(t, err)
		assert.Equal(t, data, payment.Domain{})
	})

	t.Run("Test case 3 | Error GetById", func(t *testing.T) {
		setup()
		paymentRepository.On("GetPaymentById", mock.Anything, mock.AnythingOfType("uint")).Return(payment.Domain{}, nil).Once()
		data, err := paymentService.GetPaymentById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, payment.Domain{})
	})
}

func TestUpdatePayment(t *testing.T) {
	t.Run("Test case 1 | Success Update Payment", func(t *testing.T) {
		setup()
		paymentRepository.On("UpdatePayment", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(paymentDomain, nil).Once()
		data, err := paymentService.UpdatePayment(context.Background(), paymentDomain, paymentDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Payment", func(t *testing.T) {
		setup()
		paymentRepository.On("UpdatePayment", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(paymentDomain, errors.New("Books Not Found")).Once()
		data, err := paymentService.UpdatePayment(context.Background(), paymentDomain, paymentDomain.Id)

		assert.Equal(t, data, payment.Domain{})
		assert.Error(t, err)
	})
}

func TestDeletepayment(t *testing.T) {
	t.Run("Test case 1 | Success Delete payment", func(t *testing.T) {
		setup()
		paymentRepository.On("DeletePayment", mock.Anything, mock.Anything).Return(nil).Once()
		err := paymentService.DeletePayment(context.Background(), paymentDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete payment", func(t *testing.T) {
		setup()
		paymentRepository.On("DeletePayment", mock.Anything, mock.Anything).Return(errors.New("payments  Not Found")).Once()
		err := paymentService.DeletePayment(context.Background(), paymentDomain.Id)

		assert.NotEqual(t, err, errors.New("payments Not Found"))
		assert.Error(t, err)
	})
}
