package main

import "fmt"

func main() {
	age := 10           //literal
	name := "Jack"      //literal
	rightHanded := true //literal

	fmt.Printf("%s is %d years old. Right handed %t:", name, age, rightHanded) //expresiones por que pueden ser evaluadas

	ageInTenYears := age + 10 //expresion

	fmt.Printf("%s is going to be %d years old in ten years.", name, ageInTenYears) //expresiones

	isATeenager := age >= 13 //expresion

	fmt.Printf("%s is a teenager? %t", name, isATeenager) //expresiones
}
