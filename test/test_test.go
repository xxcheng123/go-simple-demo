package main

import (
	"fmt"
	"testing"
)

var a = 1

func Test01(t *testing.T) {
	fmt.Println(a)
	var a = 2
	fmt.Println(a)
	var m = map[string]string{
		"name": "xxcheng",
	}
	for index := range m {
		fmt.Println(index)
	}
}
