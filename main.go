package main

import (
	"fmt"
	"strings"
)

// this is entry point for our main package
func main() {
	// declare a var of type int
	var m1 int
	m1 = 2

	// declare multiple vars
	// when we dont declare the int type, its automatically implicit becuase we set it equal to int type
	var (
		m2 = 4
		m3 = 6
	)

	// we cant add items of two different data types
	// we CAN typecast it and convert one data type into another
	var m4 int32
	var m5 int64

	// we can also decalre initalize vars in one line
	m6 := 85
	m7 := 22

	fmt.Println(m1)
	fmt.Println(m2 + m3)
	fmt.Println(int64(m4) + m5) // this value is typecasted
	fmt.Println(m6 + m7)

/*************************************************************/

	var n1 string
	n1 = "my name"

	n2 := "foo bar"
	n3 := "bar"

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3 + n2)
	// we can find substrings in a larger str. Check if n3 is within n2
	fmt.Println(strings.Contains(n2, n3))
	// there are lots of methods on strings to do
	fmt.Println(strings.ReplaceAll(n1, "m", "NO"))
	fmt.Println(strings.Split(n1, " ")) // returns arr with each value split by a space
}
