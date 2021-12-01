package mocks

import (
	"cleanarch/business/users"
	"context"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) Register(domain *users.Domain, ctx context.Context) (users.Domain, error) {
	ret := _m.Called(domain, ctx)
	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain, context.Context) users.Domain); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(*users.Domain, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *Repository) GetAllUser(ctx context.Context) ([]users.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []users.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
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
func (_m *Repository) Login(domain users.Domain, ctx context.Context) (users.Domain, error) {
	ret := _m.Called(domain, ctx)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(users.Domain, context.Context) users.Domain); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.Domain, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetUserById(ctx context.Context, id uint) (users.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) users.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdateUser(ctx context.Context, domain users.Domain, id uint) (users.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, users.Domain, uint) users.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, users.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) DeleteUser(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
