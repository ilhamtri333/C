package users_test

import (
	"cleanarch/app/middleware"
	"cleanarch/business/users"
	"cleanarch/business/users/mocks"
	"time"

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

// func TestRegister(t *testing.T) {
// 	setup()
// 	userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil)
// 	t.Run("Test Case 1 | Success Register", func(t *testing.T) {
// 		user, err := userService.Register(users.Domain{
// 			Id:       1,
// 			Nama:     "ilhamtw",
// 			Email:    "ilhamtw@gmail.com",
// 			Password: "123nice",
// 		}, context.Background())
// 		assert.NoError(t, err)
// 		assert.Equal(t, userDomain, user)
// 	})
// }

// func TestLogin(t *testing.T) {
// 	t.Run("Test Case 1 | Success Login", func(t *testing.T) {
// 		setup()
// 		userRepository.On("Login",
// 			mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
// 		user, token, err := userService.Login(users.Domain{
// 			Email:    "ilhamtw@gmail.com",
// 			Password: "nicenice",
// 		}, context.Background(), "123")

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

// =======================================================================================================================
// func TestLogin(t *testing.T) {

// 	setup()
// 	userRepository.On("Login",
// 		mock.Anything,
// 		mock.AnythingOfType("users.Domain")).Return(userDomain, nil).Once()

// 	t.Run("Test case 1 | Valid Login", func(t *testing.T) {
// 		user, err := userService.Login(users.Domain{
// 			Email:    "daaffa@net.usc",
// 			Password: "kecoak11",
// 		}, context.Background())

// 		assert.Nil(t, err)
// 		assert.Equal(t, users.Domain{}, user.Nama)
// 	})

// 	t.Run("Test case 2 | Error Login", func(t *testing.T) {
// 		userRepository.On("Login",
// 			mock.Anything,
// 			mock.AnythingOfType("users.Domain")).Return(users.Domain{}, errors.New("Unexpected Error")).Once()
// 		user, err := userService.Login(users.Domain{
// 			Email:    "daaffa@net.usc",
// 			Password: "kecoak11",
// 		}, context.Background())

// 		assert.NotNil(t, err)
// 		assert.Equal(t, user, users.Domain{})
// 	})

// 	t.Run("Test Case 3 | Invalid Email Empty", func(t *testing.T) {

// 		_, err := userService.Login(users.Domain{
// 			Email:    "",
// 			Password: "kecoak11",
// 		}, context.Background())
// 		assert.NotNil(t, err)
// 	})

// 	t.Run("Test Case 4 | Invalid Password Empty", func(t *testing.T) {

// 		_, err := userService.Login(users.Domain{
// 			Email:    "daaffa@net.usc",
// 			Password: "",
// 		}, context.Background())
// 		assert.NotNil(t, err)
// 	})

// }
