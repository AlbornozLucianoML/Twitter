package domain

type User struct {
	Name string
	Mail string
}

func NewUser(name string, mail string) *User {

	user := User{name, mail}

	return &user

}