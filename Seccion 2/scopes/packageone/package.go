package packageone

import "fmt"

var PackageVar = "Package variable of packageone"

func PrintMe(a, b string) {
	fmt.Println(a, b, PackageVar)
}
