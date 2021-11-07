package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "This is a clear example of why we search in one case only."

	fmt.Println(finalLetterToUpper(myString))
	fmt.Println(strings.Title(strings.ToLower(myString)))
}

func finalLetterToUpper(s string) string {
	var result string
	lenString := len(s)

	for j := 0; j < lenString; j++ {
		if strings.Contains(string(s[j]), " ") {
			result = result[:j-1] + strings.ToUpper(string(s[j-1])) + " "
		} else if strings.Contains(string(s[j]), ".") {
			result = result[:j-1] + strings.ToUpper(string(s[j-1])) + "."
		} else {
			result += string(s[j])
		}
	}

	return result

}
