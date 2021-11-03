package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 3000, FullTime: false},
	{FirstName: "Ana", LastName: "Laura", Salary: 60000, FullTime: true},
	{FirstName: "Jess", LastName: "Alexsee", Salary: 7000, FullTime: true},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	log.Println(myStaff.All())

	staff.OverPaidLimit = 2000

	log.Println(myStaff.Overpaid())

	log.Println(myStaff.UnderPaid())
}
