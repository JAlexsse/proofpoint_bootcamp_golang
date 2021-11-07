package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	go listenForKey()

	fmt.Println("Press any key or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
}

func listenForKey() {
	for {
		key := <-keyPressChan
		fmt.Println("You passed", string(key))
	}
}
