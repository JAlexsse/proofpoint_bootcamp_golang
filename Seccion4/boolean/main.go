package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	apples := 18
	oranges := 9

	//boolean expressions
	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// > < >= <=
	fmt.Printf("%d > %d: %t\n", apples, oranges, apples > oranges)
	fmt.Printf("%d < %d: %t\n", apples, oranges, apples < oranges)
	fmt.Printf("%d >= %d: %t\n", apples, oranges, apples >= oranges)
	fmt.Printf("%d <= %d: %t\n", apples, oranges, apples <= oranges)

	//compound booleans
	empl1 := Employee{
		Name:     "Jack Jones",
		Age:      34,
		Salary:   4000,
		FullTime: false,
	}

	empl2 := Employee{
		Name:     "Ana Smith",
		Age:      42,
		Salary:   7000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, empl1)
	employees = append(employees, empl2)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is older than 30.")
		} else {
			fmt.Println(x.Name, "is younger than 30.")
		}

		if x.Age > 30 && x.Salary > 4500 {
			fmt.Println(x.Name, "is older than 30 and their salary is greater than:", x.Salary)
		} else {
			fmt.Println(x.Name, "is younger than 30 or their salary is less than:", x.Salary)
		}
	}

}
