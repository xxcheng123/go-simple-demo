package main

import (
	"fmt"
	"testing"
)

func Test_0(t *testing.T) {
	var a uint = 8
	b := -1 ^ (-1 << a)
	//var c uint16 = -1 << a
	fmt.Printf("%#v,%T\n", b, b)
	//fmt.Printf("%#v,%T\n", c, c)
}
