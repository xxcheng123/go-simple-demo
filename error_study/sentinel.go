package main

import (
	"errors"
	"fmt"
)

var (
	ErrorPassIncorrect = errors.New("password incorrect")
	ErrorUserBan       = errors.New("user is baned")
)

type User struct {
	Username string
	Password string
	Status   bool
}

func main() {
	u := &User{
		Username: "xxcheng",
		Password: "123",
		Status:   false,
	}
	err := checkUser(u, "111")
	fmt.Println(err)
}

func checkUser(u *User, pass string) error {
	if u.Password != pass {
		return ErrorPassIncorrect
	} else if !u.Status {
		return ErrorUserBan
	}
	return nil
}
