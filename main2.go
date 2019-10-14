package main

import "fmt"

// arrays must be defined and symbolic of a particular data type
func todo() {
	var arr []int
	arr2 := []int{1,2,3,4}
	arr3 := []string{"hi", "my", "nameis"}

	fmt.Println(arr)
	fmt.Println(arr2, arr3)

	// append to arr
	arr3 = append(arr3, "awesome", "ok", "cool")
	fmt.Println(arr3)
}

// our func has both params of type int
// we use the dereference pointer to extract values out of the address passed in
func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

// we use pointers to store the address of another variable
//  & symbol is the referencing operator, used to get address of var
//  to get value stored inside an address, we need dereference pointer or a * symbol
func pointer() {
	m1 := "good"
	ptr := &m1
	fmt.Println(*ptr)

	// we can define like this
	m4, m5 := 400, 500
	fmt.Println(m4, m5)
	// we want to pass the pointer to a function. or the address
	swap(&m4, &m5)
	fmt.Println(m4, m5)
}

// structure- used to group logically related data. known as data encapsulation
type Car struct {
	Name string
	Age int
	ModelNo int
}
// we can define methods on top of our structures! we pass along the structure itself and define the name of our method
// this method is attached to the object of this structure
func (c Car) Print() {
	fmt.Println(c)
}
func (c Car) Drive() {
	fmt.Println("Driving....")
}
// we can return values aswell
func (c Car) GetName() string {
	return c.Name
}

// if we define other functions, we can just call them in our main()
func main() {
	// We use structures to make things off of it
	d := Car{
		Name: "bmw",
		Age: 5,
		ModelNo: 3,
	}
	d.Print()
	d.Drive()
	fmt.Println(d.GetName())

	pointer()
	// todo()
	// fmt.Println()
}
