package mocks_test

import (
	detailfas "cleanarch/business/detail_fas"
	"context"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertDetailFas(ctx context.Context, domain *detailfas.Domain) (detailfas.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 detailfas.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *detailfas.Domain) detailfas.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(detailfas.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *detailfas.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllDetailFas(ctx context.Context) ([]detailfas.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []detailfas.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []detailfas.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]detailfas.Domain)
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

func (_m *Repository) GetDetailFasById(ctx context.Context, id uint) (detailfas.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 detailfas.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) detailfas.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(detailfas.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdateDetailFas(ctx context.Context, domain detailfas.Domain, id uint) (detailfas.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 detailfas.Domain
	if rf, ok := ret.Get(0).(func(context.Context, detailfas.Domain, uint) detailfas.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(detailfas.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, detailfas.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) DeleteDetailFas(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
