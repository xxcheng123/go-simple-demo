package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := Three()
	err2 := SThree()
	fmt.Println(errors.Is(err, ErrorBase))
	fmt.Println(errors.Cause(err))
	fmt.Println(err2)
}

var (
	ErrorBase = errors.New("error base")
)

func One() error {
	return ErrorBase
}
func Two() error {
	err := One()
	return errors.Wrapf(err, "Two error")
}
func Three() error {
	err := Two()
	return errors.Wrapf(err, "Three error")
}

func SOne() error {
	return ErrorBase
}
func STwo() error {
	return errors.Errorf("%s STwo", SOne())
}
func SThree() error {
	return errors.Errorf("%s SThree", STwo())
}
