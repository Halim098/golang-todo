package Controller

import (
	"errors"
	"strings"
	helper "todo/Helper"
	"todo/Model"
	"todo/Repository"
)

type UserController interface {
	RegisterUser(data *Model.User) error
	LoginUser(data *Model.LoginInput) (string, error)
}

type userController struct {
	UserRepository Repository.UserRepository
}

func NewUserController(UserRepository Repository.UserRepository) *userController {
	return &userController{UserRepository}
}

func (u *userController) RegisterUser(data *Model.User) error {
	data.Email = strings.ToLower(data.Email)

	result, _ := u.UserRepository.GetUserByEmail(data.Email)
	if result.Email == data.Email {
		return errors.New("email already exists")
	}

	if data.Username == result.Username {
		return errors.New("username already exists")
	}

	err := u.UserRepository.BeforeSave(data)

	if err != nil {
		return err
	}

	err = u.UserRepository.SaveUser(data)

	if err != nil {
		return err
	}

	return nil
}

func (u *userController) LoginUser(data *Model.LoginInput) (string, error) {
	data.Email = strings.ToLower(data.Email)

	var token string
	user, err := u.UserRepository.GetUserByEmail(data.Email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	err = Repository.ValidatePassword(data.Password, user.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err = helper.GenerateJWT(&user)
	if err != nil {
		return "", err
	}

	return token, nil
}
