package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	f := true
	flag := &f // were pointing to the address allocated to f in mem. Using the referencing operator

	if flag == nil {
		fmt.Println("valye is nil")
	} else if *flag {
		fmt.Println("we used dereference pointer to extract value from address and it was true")
	} else {
		fmt.Println("false")
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// create an arr of type string
	arr := []string{"Mark", "luke", "enginer"}

	// we loop through array, first param is indx, second is value itself
	for i, value := range arr {
		fmt.Println("indx is ", i)
		fmt.Println("value is ",value)
	}

	// map with key of str and value of interface__or__anything!
	myMap := make(map[string]interface{})
	myMap["name"] = "mrk"
	myMap["age"] = 24

	for k,v := range myMap {
		fmt.Printf(" Key: %s and value: %v", k, v)
	}
}
