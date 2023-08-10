package main

import (
	"fmt"
	"time"
)

type MyError struct {
	ID      int
	Caller  string
	T       time.Time
	Content string
}

func (e *MyError) Error() string {
	return e.Content
}

var index = 0

func New(content string, caller string, t time.Time) error {
	er := new(MyError)
	er.ID = index
	index++
	er.Caller = caller
	er.T = t
	er.Content = content
	return er
}

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
	u2 := &User{
		Username: "jpc",
		Password: "123",
		Status:   false,
	}
	err := checkUser2(u, "111")
	err2 := checkUser2(u2, "111")
	myerr := err.(*MyError)
	myerr2 := err2.(*MyError)
	fmt.Println(myerr, ",caller:", myerr.Caller)
	fmt.Println(myerr2, ",caller:", myerr2.Caller)
	fmt.Println(myerr2 == myerr2)

}

func checkUser2(u *User, pass string) error {
	if u.Password != pass {
		return New("pass incorrect", "checkUser2", time.Time{})
	} else if !u.Status {
		return New("user baned", "checkUser2", time.Time{})
	}
	return nil
}
