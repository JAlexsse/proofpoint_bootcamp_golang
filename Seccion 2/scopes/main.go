package main

import "scopesapp/packageone"

var myVar = "Package variable of main package"

func main() {

	var blockVar = "Block variable on main function"

	packageone.PrintMe(myVar, blockVar)

}
