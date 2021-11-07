package main

import (
	"errors"
	"fmt"
	"math"
)

func main() { //go toma el orden de presedencia para realizar operaciones
	answer := 7 + 3*4
	fmt.Println("Answer:", answer)

	answer = (7 + 3) * 4
	fmt.Println("Answer:", answer)

	//multiplication
	//area = pi r 2
	var radius = 12.0

	area := math.Pi*radius + 2

	fmt.Println("Result:", area)

	//divisition
	half := 1 / 2
	fmt.Println("Result:", half) // no puede realizarlas con integers

	//cada tipo es muy exigente, y no se pueden hacer operaciones mixtas entre
	//float32 y float64 o integers.
	halfFloat := 1.0 / 2.0
	fmt.Println("Result:", halfFloat)

	//squaring (raiding to the power)
	squared := math.Pow(3.0, 2.0) //necesitamos hacerlo con floats
	fmt.Println("Squared:", squared)

	//modules
	remainder := 50 % 3
	fmt.Println("Remainder:", remainder)

	//unary operators
	x := 3
	x++ //suma uno
	fmt.Println("Unary operator:", x)

	//PRECEDENCE
	//multiplication and division
	a := 12.0 * 3.0 / 4.0
	b := (12.0 * 3.0) / 4.0
	c := 12.0 * (3.0 / 4.0)

	fmt.Println("a:", a, "- b:", b, "- c:", c) //son todos iguales

	// integer division
	unclear := 12 * (3 / 4)
	fmt.Println("unclear:", unclear) //da cero por que esta dividiendo integers

	//parentesis, tienen la presedencia mas alta
	f := 12.0 / 3.0 / 4.0
	fmt.Println("f:", f)
	f = 12.0 / (3.0 / 4.0)
	fmt.Println("f:", f)

	//additions and subtraction
	w := 12 + 3 - 4
	y := (12 + 3) - 4
	z := 12 + (3 - 4)
	fmt.Println("w:", w, "y:", y, "z:", z) // da igual

	//does one number divide another
	x1 := 12
	y1 := 3

	if x1%y1 == 0 {
		fmt.Println("yes!")
	} else {
		fmt.Println("no!")
	}

	//relational and conditionals
	second := 31
	minute := 1

	if minute < 59 && second+1 > 59 {
		minute++
	}

	fmt.Println(minute)

	//shortcircuit evaluation
	num1 := 12
	num2 := 0

	num3, err := divideTwoNumbers(num1, num2)
	if err != nil {
		fmt.Println("error")
	} else {
		if num3 == 2 {
			fmt.Println("We found 2")
		}
	}

	//assigment operators
	j := 12
	j++
	fmt.Println(j)
	k := 10
	k *= 2
	fmt.Println(k)
	k /= 2
	fmt.Println(k)

}

func divideTwoNumbers(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}
