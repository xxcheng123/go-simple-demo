package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

// TestIPV42Num IPV4转数字地址
func TestIPV42Num(t *testing.T) {
	ipStr := "192.168.0.1"
	ipStrSlice := strings.Split(ipStr, ".")
	if len(ipStrSlice) != 4 {
		panic("ip format error.")
	}
	var IP uint32 = 0
	var _ = IP
	for index, perIPStr := range ipStrSlice {
		perIP, _ := strconv.ParseInt(perIPStr, 10, 64)
		if perIP < 0 || perIP > 255 {
			panic("ip format error.")
		}
		offset := 8 * (4 - index - 1)
		IP = IP | uint32(perIP)<<offset
		fmt.Println(perIP, index)
	}
	fmt.Println(IP)
}

// TestNum2IPV4 数字转IPV4
func TestNum2IPV4(t *testing.T) {
	var num uint32 = 665676686
	if num < 0 || num > math.MaxUint32 {
		panic("ip format error.")
	}
	ipSlice := make([]uint32, 4)
	ipSlice[0] = num >> 24 & 0xff
	ipSlice[1] = num >> 16 & 0xff
	ipSlice[2] = num >> 8 & 0xff
	ipSlice[3] = num >> 0 & 0xff
	ipStr := fmt.Sprintf("%d.%d.%d.%d", ipSlice[0], ipSlice[1], ipSlice[2], ipSlice[3])
	fmt.Println(ipStr)
}
