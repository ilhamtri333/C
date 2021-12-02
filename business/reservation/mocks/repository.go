package mocks_test

import (
	"cleanarch/business/reservation"
	"context"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertReservation(ctx context.Context, domain *reservation.Domain) (reservation.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 reservation.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *reservation.Domain) reservation.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(reservation.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *reservation.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllReservation(ctx context.Context) ([]reservation.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []reservation.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []reservation.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservation.Domain)
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

func (_m *Repository) GetReservationById(ctx context.Context, id uint) (reservation.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 reservation.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) reservation.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(reservation.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdateReservation(ctx context.Context, domain reservation.Domain, id uint) (reservation.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 reservation.Domain
	if rf, ok := ret.Get(0).(func(context.Context, reservation.Domain, uint) reservation.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(reservation.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, reservation.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) DeleteReservation(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
