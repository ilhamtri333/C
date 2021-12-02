package mocks_test

import (
	payment "cleanarch/business/payment_method"
	"context"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertPayment(ctx context.Context, domain *payment.Domain) (payment.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 payment.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *payment.Domain) payment.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(payment.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *payment.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllPayment(ctx context.Context) ([]payment.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []payment.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []payment.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]payment.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetPaymentById(ctx context.Context, id uint) (payment.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 payment.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) payment.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(payment.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdatePayment(ctx context.Context, domain payment.Domain, id uint) (payment.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 payment.Domain
	if rf, ok := ret.Get(0).(func(context.Context, payment.Domain, uint) payment.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(payment.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, payment.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) DeletePayment(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
