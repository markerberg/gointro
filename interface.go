package main

import "fmt"

func main() {
	// interface lets us enforce style and we can use it as a wildcard for unknown types

	// this means we set a map where the key is a string and the value is anything we want
	mymap := make(map[string]interface{})
	mymap["name"] = "M4rcK_231"
	mymap["age"] = 24

	fmt.Println(mymap);
}
