package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const ready = "and press ENTER when ready."

func main() {

	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	display(firstNumber, secondNumber, subtraction, answer)
}

func display(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)

	//display a welcome/instruction
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", ready)

	reader.ReadString('\n')

	//take them through the games
	fmt.Println("Multiply your number by", firstNumber, ready)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, ready)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", ready)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, ready)
	reader.ReadString('\n')

	//give them the answer
	fmt.Println("The answer is", answer)
}
