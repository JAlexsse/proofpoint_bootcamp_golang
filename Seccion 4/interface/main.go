package main

import "fmt"

//interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

//type dog
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

//type cat
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	Riddle(&dog)

	cat := Cat{
		Name:         "Kiki",
		Sound:        "meow",
		NumberOfLegs: 4,
		HasTail:      true,
	}

	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf("This animal has %d legs, and makes %s sounds.", a.HowManyLegs(), a.Says())
	fmt.Println(riddle)
}
