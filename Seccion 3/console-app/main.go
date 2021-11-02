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

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffes[i]))
	}
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
