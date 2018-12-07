package domain

import "testing"

func TestUserCreation(t *testing.T) {

	user := NewUser("", "")

	if user == nil {
		t.Error("Expected an user created, got a nil user")
	}

}

func TestUserWithName(t *testing.T) {

	name := "Luciano"

	user := NewUser(name, "")

	if user.Name != name {
		t.Errorf("Expected name = %s, got %s", user.Name, name)
	}

}

func TestUserWithMail(t *testing.T) {

	mail := "luciano@gmail.com"

	user := NewUser("", mail)

	if user.Mail != mail {
		t.Errorf("Expected mail = %s, got %s", user.Mail, mail)
	}

}

func TestUserMailIsValid(t *testing.T) {

	TestUserWithMail(t)

	

}

