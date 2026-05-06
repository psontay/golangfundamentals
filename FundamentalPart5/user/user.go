package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	dob       string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email string, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email or password is empty")
	}
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			dob:       "220605",
			createdAt: time.Now(),
		},
	}, nil
}

func New(firstName string, lastName string, dob string) (*User, error) {
	if firstName == "" || lastName == "" || dob == "" {
		return nil, fmt.Errorf("invalid input")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		dob:       dob,
		createdAt: time.Now(),
	}, nil
}

func (u User) GetUserDetail() {
	fmt.Println(u.firstName)
	fmt.Println(u.lastName)
	fmt.Println(u.dob)
	fmt.Println(u.createdAt)
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}
