package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	courseName := "Learn Go From Scratch"

	fmt.Println(courseName[0])         // asi imprime el valor de la rune que esta en el primer lugar en int
	fmt.Println(string(courseName[0])) //ahora imprime L

	fmt.Println(courseName[6])         // asi imprime el valor de la rune que esta en el primer lugar en int
	fmt.Println(string(courseName[6])) //ahora imprime G

	lenCourseName := len(courseName)

	fmt.Println("The len of the course name is", lenCourseName)

	for i := 0; i < lenCourseName; i++ {
		fmt.Print(string(courseName[i]))
	}

	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")
	mySlice = append(mySlice, "four")

	fmt.Println("mySlice has", len(mySlice), "elements")
	fmt.Println("The last element of mySlice is", mySlice[len(mySlice)-1])

	courses := []string{
		"Learn Go",
		"Learn JAVA",
		"Learn Python",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in ->", x)
			fmt.Println("Go is found in index ->", strings.Index(x, "Go"))
		}
	}

	goCourse := "Go is a great programming lenguage. Go for it!"

	fmt.Println(strings.HasPrefix(goCourse, "Go"))
	fmt.Println(strings.HasPrefix(goCourse, "Python"))

	fmt.Println(strings.HasSuffix(goCourse, "!"))

	fmt.Println(strings.Count(goCourse, "Go"))
	fmt.Println(strings.Count(goCourse, "Python"))

	fmt.Println(strings.Index(goCourse, "Go"))
	fmt.Println(strings.Index(goCourse, "Python"))

	fmt.Println(strings.LastIndex(goCourse, "Go"))
	fmt.Println(strings.LastIndex(goCourse, "Python"))

	goCourse = strings.Replace(goCourse, "Go", "Golang", 1) //reemplaza la primera ocurrencia, -1 reemplaza todas
	//goCourse = strings.ReplaceAll(goCourse, "Go", "Golang") hace lo mismo que Replace con -1

	fmt.Println(goCourse)

	//strings comparison

	badEmail := " example@something.com  "
	badEmail = strings.TrimSpace(badEmail)

	fmt.Printf("=%s=\n", badEmail)

	anyString := "something something something something"

	replacedString, err := replaceNth(anyString, "something", "another", 3)

	if err != nil {
		log.Println("err:", err)
	}

	fmt.Println(replacedString)

	//no existe en el index pedido
	replacedString2, err2 := replaceNth(anyString, "something", "another", 5)

	if err2 != nil {
		log.Println("err2:", err2)
	}

	fmt.Println(replacedString2)

	//no existe
	replacedString3, err3 := replaceNth(anyString, "found", "another", 3)

	if err3 != nil {
		log.Println("err3:", err2)
	}

	fmt.Println(replacedString3)

	//case
	myString := "This is a clear example of why we search in one case only."

	searchString := strings.ToLower(myString)

	if strings.Contains(searchString, "this") {
		fmt.Println("We could find it!")
	} else {
		fmt.Println("We could not find it! :(")
	}

	//other case functions
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
	fmt.Println(strings.Title(myString))
}

func replaceNth(s, old, new string, n int) (string, error) {
	// index
	i := 0

	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			//no lo encontro
			break
		}

		i = i + x
		if j == n {
			return s[:i] + new + s[i+len(old):], nil
		}

		i += len(old)
	}

	err := errors.New("the substring could not be found")
	return s, err
}
