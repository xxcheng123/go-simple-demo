package main

import "errors"

var (
	errorPassIncorrect = errors.New("password incorrect")
	errorUserBan       = errors.New("user is baned")
)

func IsErrorPassIncorrect(err error) bool {
	return err == errorPassIncorrect
}
func IsErrorUserBan(err error) bool {
	return err == errorUserBan
}
