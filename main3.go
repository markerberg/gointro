package main

import "fmt"

// interfaces are useful. If any stuctures follow an interface, you know which methods and properties they already have
// good way to make sure our functions follow a set of instructions we set
type Car interface {
	Drive()
	Stop()
}

// we will make lambo follow the car interface
type Lambo struct {
	LamboModel string
}

// we return the lambo model which needs to satisfy the Car interface
func NewModel(arg string) Car {
	return &Lambo{arg}
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Drive() {
	fmt.Println("moving the lambo")
	fmt.Println(l.LamboModel)
}

func (l *Lambo) Stop() {
	fmt.Println("stopping lambo")
}

func (c *Chevy) Drive() {
	fmt.Println("chevy is moving")
	fmt.Println(c.ChevyModel)
}

func main() {
	l := NewModel("gallardo")
	c := Chevy{"volt"}

	l.Drive()
	c.Drive()
}
