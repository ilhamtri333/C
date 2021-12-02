package users_test

import (
	"cleanarch/app/middleware"
	"cleanarch/business/users"
	"cleanarch/business/users/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = mocks.Repository{Mock: mock.Mock{}}
var userService users.UserUseCaseInterface
var token string
var userDomain users.Domain
var AllUserDomain []users.Domain

func setup() {
	userService = users.NewUseCase(&userRepository, time.Hour*10, &middleware.ConfigJWT{})
	userDomain = users.Domain{
		Id:       1,
		Nama:     "ilhamtw",
		Email:    "ilhamtw@gmail.com",
		Password: "123nice",
	}
	token = "tokennnnn"
	AllUserDomain = append(AllUserDomain, userDomain)
}

func TestRegister(t *testing.T) {
	setup()
	userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil)
	t.Run("Test Case 1 | Success Register", func(t *testing.T) {
		user, err := userService.Register(&users.Domain{
			Id:       1,
			Nama:     "ilhamtw",
			Email:    "ilhamtw@gmail.com",
			Password: "123nice",
		}, context.Background())
		assert.NoError(t, err)
		assert.Equal(t, userDomain, user)
	})
}

// func TestLogin(t *testing.T) {
// 	t.Run("Test Case 1 | Success Login", func(t *testing.T) {
// 		setup()
// 		userRepository.On("Login",
// 			mock.Anything, mock.Anything).Return(users.Domain{}, nil).Once()
// 		user, token, err := userService.Login(users.Domain{
// 			Email:    "ilhamtw@gmail.com",
// 			Password: "nicenice",
// 		}, context.Background())

// 		assert.NotNil(t, token)
// 		assert.NoError(t, err)
// 		assert.Equal(t, user, users.Domain{})
// 	})

// 	t.Run("Test Case 2 | Cannot Login (Password Not Found)", func(t *testing.T) {
// 		data, token, err := userService.Login(userDomain, context.Background())

// 		assert.Equal(t, users.Domain{}, data)
// 		assert.Error(t, err)
// 		assert.Equal(t, token, "")
// 	})

// 	t.Run("Test Case 3 | Cannot Login (Wrong Auth)", func(t *testing.T) {
// 		setup()
// 		userRepository.On("Login",
// 			mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("Users Not Found")).Once()
// 		data, token, err := userService.Login(userDomain.Email, context.Background())

// 		assert.Equal(t, users.Domain{}, data, token, "")
// 		assert.NoError(t, err)
// 	})
// }

func TestUser(t *testing.T) {
	t.Run("Test case 1 | Success GetAllUser", func(t *testing.T) {
		setup()
		userRepository.On("GetAllUser", mock.Anything, mock.Anything).Return(AllUserDomain, nil).Once()
		data, err := userService.GetAllUser(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(AllUserDomain))
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetById", func(t *testing.T) {
		setup()
		userRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.GetUserById(context.Background(), userDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetById(user ID = 0)", func(t *testing.T) {
		setup()
		userDomain.Id = 0
		userRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.GetUserById(context.Background(), userDomain.Id)

		assert.NoError(t, err)
		assert.Equal(t, data, users.Domain{})
	})

	t.Run("Test case 3 | Error GetById", func(t *testing.T) {
		setup()
		userRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("uint")).Return(users.Domain{}, nil).Once()
		data, err := userService.GetUserById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, users.Domain{})
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Test case 1 | Success Update user", func(t *testing.T) {
		setup()
		userRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.UpdateUser(context.Background(), userDomain, userDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update user", func(t *testing.T) {
		setup()
		userRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, errors.New("Books Not Found")).Once()
		data, err := userService.UpdateUser(context.Background(), userDomain, userDomain.Id)

		assert.Equal(t, data, users.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Test case 1 | Success Delete user", func(t *testing.T) {
		setup()
		userRepository.On("DeleteUser", mock.Anything, mock.Anything).Return(nil).Once()
		err := userService.DeleteUser(context.Background(), userDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete user", func(t *testing.T) {
		setup()
		userRepository.On("DeleteUser", mock.Anything, mock.Anything).Return(errors.New("users  Not Found")).Once()
		err := userService.DeleteUser(context.Background(), userDomain.Id)

		assert.NotEqual(t, err, errors.New("users Not Found"))
		assert.Error(t, err)
	})
}
