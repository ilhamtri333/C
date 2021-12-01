package request

import "cleanarch/business/users"

type UserRegister struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		Nama:     user.Nama,
		Email:    user.Email,
		Password: user.Password,
	}
}
