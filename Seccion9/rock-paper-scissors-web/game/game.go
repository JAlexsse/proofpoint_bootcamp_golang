package game

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
	WIN      = 1
	LOST     = 2
	DRAW     = 3
)

var winMessages = []string{
	"Good job!",
	"Luck had been on your side!",
	"Nice!",
	"Keep going!",
}

var loseMessages = []string{
	"Don't give up!",
	"Oh, dang!",
	"You will have better luck next time!",
	"Cheer up!",
}

var drawMessages = []string{
	"Masterminds think the same...",
	"Impressive!",
	"Who will win next?",
	"Interesting, let's keep playing!",
}

type Round struct {
	Winner         int    `json:"winner"`
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(4)
	computerChoice := ""
	roundResult := ""
	winner := 0
	message := ""

	switch computerValue {
	case ROCK:
		computerChoice = "ROCK"
	case PAPER:
		computerChoice = "PAPER"
	case SCISSORS:
		computerChoice = "SCISSORS"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
		message = messageSelect(winner)
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "You win!"
		winner = WIN
		message = messageSelect(winner)
	} else {
		roundResult = "Oh no.. You lost! :("
		winner = LOST
		message = messageSelect(winner)
	}

	result := Round{
		Winner:         winner,
		ComputerChoice: computerChoice,
		Message:        message,
		RoundResult:    roundResult,
	}

	return result
}

func messageSelect(n int) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(4)
	var message string
	switch n {
	case 1:
		message = winMessages[index]
	case 2:
		message = loseMessages[index]
	case 3:
		message = drawMessages[index]
	default:
	}
	return message
}
