package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsACat        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	values := make(map[rune]bool)
	values['y'] = true
	values['n'] = false

	var user User
	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your favourite number?")
	user.OwnsACat = values[readCats("Do you have a cat (y/n)?")]
	fmt.Printf("Your name is %s.\nYou are %d years old.\nYour favourite number is %.2f\nYou own a cat: %t",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		user.OwnsACat)
}

func prompt() {
	fmt.Println("->")
}

func readString(question string) string {
	for {
		fmt.Println(question)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		if userInput == "" {
			fmt.Println("Please enter a value.")
		} else {
			return userInput
		}
	}
}

func readInt(question string) int {
	for {
		fmt.Println(question)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}

	}
}

func readFloat(question string) float64 {
	for {
		fmt.Println(question)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64) //para parsear float tenemos que especificar si 32 o 64
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}

	}
}

func readCats(question string) rune {

	//abre teclado
	err := keyboard.Open()

	//error de apertura de teclado
	if err != nil { //si es diferente a nada
		log.Fatal(err)
	}

	//prepara la funcion para cerrar el teclado
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		//informa pregunta
		fmt.Println(question)
		prompt()

		//capta la tecla
		char, _, err := keyboard.GetSingleKey()

		//error de captacion de tecla
		if err != nil {
			log.Fatal(err)
		}

		//validacion de char
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else {
			return char
		}

	}
}
