package mocks_test

import (
	"cleanarch/business/kost"
	"context"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertKost(ctx context.Context, domain *kost.Domain) (kost.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 kost.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *kost.Domain) kost.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(kost.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *kost.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllKost(ctx context.Context) ([]kost.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []kost.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []kost.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kost.Domain)
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

func (_m *Repository) GetKostById(ctx context.Context, id uint) (kost.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 kost.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) kost.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(kost.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdateKost(ctx context.Context, domain kost.Domain, id uint) (kost.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 kost.Domain
	if rf, ok := ret.Get(0).(func(context.Context, kost.Domain, uint) kost.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(kost.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kost.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) DeleteKost(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
