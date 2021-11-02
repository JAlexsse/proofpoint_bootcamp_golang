package main

import (
	"fmt"
	"log"
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

func main() {
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

	myFirstPointer := &myInt
	fmt.Println("The value of myInt is on the space of memory:", myFirstPointer)
	*myFirstPointer = 14
	fmt.Println("Now, the value of myInt is:", myInt)

	changeValueOfPointer(&myInt)

	fmt.Println("Now the value of myInt is:", myInt)

}

func changeValueOfPointer(num *int) {
	*num = 25
}
