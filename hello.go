package main

import (
	"fmt"
	"strconv"
)

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
	fmt.Printf(strconv.Itoa(aa))

}
