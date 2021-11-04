package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
	WON      = 3
	LOST     = 4
	DRAW     = 5
)

func main() {

	playerCounter := 0
	computerCounter := 0
	drawCounter := 0

	clearScreen()

	fmt.Println("------")
	fmt.Println("Play ROCK-PAPER-SCISSORS with the computer! You have 3 rounds... Let's play!")
	fmt.Println("------")

	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		computerValue := rand.Intn(3)

		playerValue := playerTurn()

		result := play(playerValue, computerValue)

		fmt.Println("------")

		switch result {
		case WON:
			playerCounter++
		case LOST:
			computerCounter++
		case DRAW:
			drawCounter++
		}
	}

	if playerCounter == computerCounter || drawCounter == 3 {
		fmt.Println("It's a draw!")
	} else if playerCounter > computerCounter {
		fmt.Println("You won the game!")
	} else {
		fmt.Println("Oh no... You lost!")
	}
	fmt.Println("------")

	fmt.Printf("Score: Player (%d/3) Computer (%d/3) Draws: (%d/3)\n", playerCounter, computerCounter, drawCounter)

	fmt.Println("------")

}

func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		//linux
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func play(playerValue, computerValue int) int {

	result := -1

	//switch
	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose ROCK")
	case PAPER:
		fmt.Println("Computer chose PAPER")
	case SCISSORS:
		fmt.Println("Computer chose SCISSORS")
	}

	fmt.Println("------")
	fmt.Println("WINNER")
	fmt.Println("------")

	if playerValue == computerValue {
		fmt.Println("It's a draw!")
		result = DRAW
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				fmt.Println("You lose!")
				result = LOST
			} else {
				fmt.Println("You win!")
				result = WON
			}
		case PAPER:
			if computerValue == SCISSORS {
				fmt.Println("You lose!")
				result = LOST
			} else {
				fmt.Println("You win!")
				result = WON
			}
		case SCISSORS:
			if computerValue == ROCK {
				fmt.Println("You lose!")
				result = LOST
			} else {
				fmt.Println("You win!")
				result = WON
			}
		default:
			fmt.Println("Invalid choice.")
		}
	}
	return result
}

func playerTurn() int {
	playerChoice := ""
	playerValue := -1
	validValue := false

	reader := bufio.NewReader(os.Stdin)

	for !validValue {
		fmt.Print("Please enter rock, paper or scissors ->")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)

		switch playerChoice {
		case "rock":
			playerValue = ROCK
			validValue = true
		case "paper":
			playerValue = PAPER
			validValue = true
		case "scissors":
			playerValue = SCISSORS
			validValue = true
		default:
			fmt.Println("Invalid choice! Try again!")
		}
	}
	fmt.Printf("You chose %s\n", playerChoice)
	return playerValue
}
