package mocks_test

import (
	"cleanarch/business/owner"
	"context"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertOwner(ctx context.Context, domain *owner.Domain) (owner.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 owner.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *owner.Domain) owner.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(owner.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *owner.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllOwner(ctx context.Context) ([]owner.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []owner.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []owner.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]owner.Domain)
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

func (_m *Repository) GetOwnerById(ctx context.Context, id uint) (owner.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 owner.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) owner.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(owner.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdateOwner(ctx context.Context, domain owner.Domain, id uint) (owner.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 owner.Domain
	if rf, ok := ret.Get(0).(func(context.Context, owner.Domain, uint) owner.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(owner.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, owner.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) DeleteOwner(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
