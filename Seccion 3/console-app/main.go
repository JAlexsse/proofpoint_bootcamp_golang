package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	if err != nil { //si es diferente a nada
		log.Fatal(err)
	}

	//defer se ejecutara inmediatamente despues de que
	//la funcion finalice
	defer func() {
		_ = keyboard.Close()
	}()

	coffes := make(map[int]string)
	coffes[1] = "Cappucino"
	coffes[2] = "Late"
	coffes[3] = "Americano"
	coffes[4] = "Mocca"
	coffes[5] = "Macchiato"
	coffes[6] = "Espresso"

	menu()

	//section5 arrangment
	char := ' ' // no puede ser '' por que no es un rune, al menos tiene que ser un caracter

	for char != 'q' && char != 'Q' {
		char, _, err = keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char)) //transformamos char en int
		el, ok := coffes[i]                //buscamos en el map si existe el valor
		if ok {                            //si existe el valor entonces lo imprime
			fmt.Printf("You chose %s\n", el)
		}
	}

	fmt.Println("Thanks for using our services.")
}

func menu() {
	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1. Cappucchino")
	fmt.Println("2. Latte")
	fmt.Println("3. Americano")
	fmt.Println("4. Mocca")
	fmt.Println("5. Macchiato")
	fmt.Println("6. Espresso")
	fmt.Println("Q. QUIT")
}
