package staff

import "fmt"

var OverPaidLimit = 3000  //publica
var underPaidLimit = 3000 // privada

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee
	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

func (e *Office) UnderPaid() []Employee {
	var underpaid []Employee
	for _, x := range e.AllStaff {
		if x.Salary < underPaidLimit {
			underpaid = append(underpaid, x)
		}
	}
	return underpaid
}

func (e *Office) notVisible() { // function privada
	fmt.Println("Hello, World!")
}

func myFunction() { //function privada sin reciver, no ve los parametros de Office ni los de Employee
	fmt.Println("I'm a function")
}
