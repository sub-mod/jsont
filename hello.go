package main

import (
	"fmt"
	"strconv"
	"unicode"
)


//https://github.com/golang/go/wiki/GoForCPPProgrammers
func main() {
	fmt.Println("example ")

	//variable
	a := 10

	//constant
	const IPv4Len = 4
	const (
		aa = 100
	)
	fmt.Println(a)

	//Convert string to integer
	fmt.Println(strconv.Itoa(aa))

	fmt.Println(unicode.IsUpper('A'))
	fmt.Println(unicode.IsLower('a'))

	//https://github.com/golang/go/wiki/Switch
}

type Circle struct {
	x float64
	y float64
	r float64
}

const (
	One Type = iota // 1
	Two
	Three
)

type Type int

func (t Type) String() string {
	s := ""
	if t&One == One {
		s += "One"
	}
	return s
}
