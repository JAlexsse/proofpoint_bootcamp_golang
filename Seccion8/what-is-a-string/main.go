package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello world!"
	fmt.Println("String:", name)
	fmt.Println()

	fmt.Println("Bytes")

	for i := 0; i < len(name); i++ {
		fmt.Printf("%x", name[i])
	}
	fmt.Println()

	fmt.Println("Index\tRune\tString") //como puede contener runas puede contener cualquier clase de caracteres
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println("Three ways to concatenate strings")
	h := "hello, "
	w := "world!"

	//1
	myString := h + w
	fmt.Println(myString)

	//2
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)

	// using stringbuilder
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println((sb.String()))

	fmt.Println()
	name = "AJDLASJDAJLSJDLAKJSLDJLASAJS"
	fmt.Print(name[1:13])

}
