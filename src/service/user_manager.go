package service

import "github.com/AlbornozLucianoML/Twitter/src/domain"

type UserManager struct {
	UsersRegistered []*domain.User
	UsersLoggedIn []*domain.User
}

func NewUserManager() *UserManager {

	usersRegistered := make([]*domain.User, 0)
	usersLoggedIn := make([]*domain.User, 0)

	userManager := UserManager{usersRegistered, usersLoggedIn}

	return &userManager

}

func (userManager *UserManager) RegisterUser(user *domain.User) {

	userManager.UsersRegistered = append(userManager.UsersRegistered, user)

}