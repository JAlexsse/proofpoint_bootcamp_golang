package main

import (
	"fmt"
	"log"
	"sort"
)

//basic types

var myInt int

var myUint uint //solo puede tener valores positivos o 0

var myFloat float32
var myFloat64 float64 //puede contener numeros mucho mas grandes

var myString string

var myBool bool

//aggregate types

type myStruct struct { //struct
	NumberOfStruct int
	StringOfStruct string
	BoolOfStruct   bool
}

//reference types

//interface type

type MyType struct {
	Name   string
	Style  string
	Number int
}

func (a *MyType) Says() { //como declarar una funcion para un type
	fmt.Printf("MyType name is %s, has style %s and his number is %d\n", a.Name, a.Style, a.Number)
}

func (a *MyType) numberOfType() {
	fmt.Println("The number of mytype is:", a.Number)
}

func main() {
	//go rutin
	go doSomething("Hello, go rutin!")

	//basic types
	myInt = 10
	myUint = 20
	myFloat = 10.1
	myFloat64 = 100.1
	myString = "String"
	myBool = true

	log.Println(myInt, myUint, myFloat, myFloat64, myString, myBool)

	//array
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "horse"

	//struct
	newStruct := myStruct{
		BoolOfStruct:   true,
		NumberOfStruct: 2,
		StringOfStruct: "It's a struct",
	}

	fmt.Printf("My struct is %s, %d, %t\n",
		newStruct.StringOfStruct,
		newStruct.NumberOfStruct,
		newStruct.BoolOfStruct)

	//pointer
	myFirstPointer := &myInt
	fmt.Println("The value of myInt is on the space of memory:", myFirstPointer)
	*myFirstPointer = 14
	fmt.Println("Now, the value of myInt is:", myInt)

	changeValueOfPointer(&myInt)

	fmt.Println("Now the value of myInt is:", myInt)

	//slices
	var mySlices []string
	mySlices = append(mySlices, "slice1")
	mySlices = append(mySlices, "slice3")
	mySlices = append(mySlices, "slice2")
	mySlices = append(mySlices, "slice4")

	fmt.Println(mySlices)

	for i, x := range mySlices { //forma de recorrer slices o array
		fmt.Println(i, x)
	}

	fmt.Println("First element:", mySlices[0]) //como leer un slice
	fmt.Println("First two elements:", mySlices[0:2])
	fmt.Println("MySlices has", len(mySlices), "elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(mySlices)) //SORT SLICES
	sort.Strings(mySlices)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(mySlices))

	mySlices = deleteFromSlices(mySlices, 1) //DELETE ONE ELEMENT, rompe con el sort y hay que realizarlo de nuevo

	sort.Strings(mySlices)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(mySlices))

	//maps
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	myMap["four"] = 4
	myMap["five"] = 5

	for key, value := range myMap { //como leer map
		fmt.Println(key, value)
	}

	delete(myMap, "four") //como eliminar

	el, ok := myMap["four"] //como chequear que un valor esta en el map
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println("Is not in map") //si no lo encuentra deja el en su valor default, en este caso 0
	}

	myMap["two"] = 4 //como escribir en mapa

	//functions
	fmt.Println(addTwoNumber(2, 4))
	fmt.Println(sumMany(2, 4, 5, 6, 7, 8))

	//functions in types
	var newType MyType
	newType.Name = "Type name"
	newType.Style = "Type style"
	newType.Number = 123

	newType.Says()

	newType2 := MyType{
		Name:   "Type name 2",
		Style:  "Type style 2",
		Number: 12345,
	}

	newType2.Says()
	newType2.numberOfType()

}

func changeValueOfPointer(num *int) {
	*num = 25
}

func deleteFromSlices(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}

func addTwoNumber(x, y int) int {
	return x + y
}

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums { //ignora indice
		total = total + x
	}

	return total
}

//go rutin
func doSomething(s string) {
	until := 0
	for {
		fmt.Println("s is:", s)

		until := until + 5

		if until >= 5 {
			break
		}
	}
}
