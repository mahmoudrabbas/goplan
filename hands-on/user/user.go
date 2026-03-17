package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	userFirstName string
	userLastName  string
	userDate      string
	// age int
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			userFirstName: "Mahmoud",
			userLastName:  "Ramdan",
			userDate:      "---",
			createdAt:     time.Now(),
		},
	}
}

func New(firstName, lastName, date string) (*User, error) {
	if firstName == "" || lastName == "" || date == "" {
		return nil, errors.New("FirstName, LastName, Date Are Required..\n")
	}

	return &User{
		userFirstName: firstName,
		userLastName:  lastName,
		userDate:      date,
	}, nil
}

func (u *User) OutputUserData() {
	format := "FirstName: %v\nLastName: %v\nDate: %v\n"
	fmt.Printf(format, u.userFirstName, u.userLastName, u.userDate)
}

func (u *User) ClearUserName() {
	u.userFirstName = ""
	u.userLastName = ""
}
