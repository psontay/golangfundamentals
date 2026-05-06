package practice

import "errors"

type User struct {
	ID    int
	Name  string
	Email string
}

func FindUser(id int) (*User, error) {
	if id == 1 {
		return &User{
			ID:    1,
			Name:  "Son Tay",
			Email: "sontay",
		}, nil
	}
	return nil, errors.New("user not found")
}
