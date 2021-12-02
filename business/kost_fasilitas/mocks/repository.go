package mocks_test

import (
	kostfasilitas "cleanarch/business/kost_fasilitas"
	"context"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertKostFas(ctx context.Context, domain *kostfasilitas.Domain) (kostfasilitas.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 kostfasilitas.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *kostfasilitas.Domain) kostfasilitas.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(kostfasilitas.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *kostfasilitas.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllKostFas(ctx context.Context) ([]kostfasilitas.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []kostfasilitas.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []kostfasilitas.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kostfasilitas.Domain)
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

func (_m *Repository) GetKostFasById(ctx context.Context, id uint) (kostfasilitas.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 kostfasilitas.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) kostfasilitas.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(kostfasilitas.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdateKostFas(ctx context.Context, domain kostfasilitas.Domain, id uint) (kostfasilitas.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 kostfasilitas.Domain
	if rf, ok := ret.Get(0).(func(context.Context, kostfasilitas.Domain, uint) kostfasilitas.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(kostfasilitas.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kostfasilitas.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) DeleteKostFas(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
