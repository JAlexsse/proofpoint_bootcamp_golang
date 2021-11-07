package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle //embeber el type, tiene acceso a Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passengers:", v.NumberOfPassengers)
	fmt.Println("Number of wheels:", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is Electric:", c.IsElectric)
	fmt.Println("Is Hybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {
	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 2,
	}

	volvo := Car{
		Make:       "Volvo",
		Model:      "XC90 T8",
		Year:       2021,
		IsElectric: false, // si no se especifica es falso
		IsHybrid:   true,
		Vehicle:    suv,
	}

	volvo.show()
}
