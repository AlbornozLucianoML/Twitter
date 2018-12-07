package service

import (
	"github.com/AlbornozLucianoML/Twitter/src/domain"
	"testing"
)

func TestUserManagerCreation(t *testing.T) {

	userManager := NewUserManager()

	if (userManager == nil) {
		t.Error("User manager is nil")
	}

}

func TestUserManagerWithUsersListEmpty(t *testing.T) {

	userManager := NewUserManager()

	if userManager.UsersRegistered == nil {
		t.Error("Registered user list is nil")
	}

	if userManager.UsersLoggedIn == nil {
		t.Error("Logged in user list is nil")
	}

}

func TestUserManagerRegisterUser(t *testing.T) {

	userManager := NewUserManager()
	user := domain.NewUser()

	userManager.RegisterUser(user)

	if len(userManager.UsersRegistered) != 1 {
		t.Error("Expected 1 user registered, got: ", len(userManager.UsersRegistered))
	}

}
