package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Grid struct {
	Position1 int `json:"position_1"`
	Position2 int `json:"position_2"`
	Position3 int `json:"position_3"`
	Position4 int `json:"position_4"`
	Position5 int `json:"position_5"`
	Position6 int `json:"position_6"`
	Position7 int `json:"position_7"`
	Position8 int `json:"position_8"`
	Position9 int `json:"position_9"`
}

type Round struct {
	Grid    Grid   `json:"grid"`
	Turn    int    `json:"turn"`
	Winner  int    `json:"winner"`
	IsDraw  bool   `json:"is_draw"`
	Message string `json:"message"`
}

var playingMessages = []string{
	"Good job!",
	"Keep going!",
	"Interesting, let's keep playing!",
}

func PlayRound(playerValue int, round *Round) {
	round.Turn++
	player := 0

	if round.Turn%2 != 0 {
		player = 1
	} else {
		player = 2
	}

	switch playerValue {
	case 1:
		round.Grid.Position1 = player
	case 2:
		round.Grid.Position2 = player
	case 3:
		round.Grid.Position3 = player
	case 4:
		round.Grid.Position4 = player
	case 5:
		round.Grid.Position5 = player
	case 6:
		round.Grid.Position6 = player
	case 7:
		round.Grid.Position7 = player
	case 8:
		round.Grid.Position8 = player
	case 9:
		round.Grid.Position9 = player
	}

	round.Winner = thereIsAWinner(player, round)

	switch round.Winner {
	case 0:
		if round.Turn == 9 {
			round.Message = "It's a draw!"
			round.IsDraw = true
		} else {
			round.Message = messageSelect()
		}

	default:
		round.Message = fmt.Sprintf("The winner is player %d", round.Winner)
	}

}

func thereIsAWinner(player int, round *Round) int {

	winner := 0

	if round.Grid.Position1 == player && round.Grid.Position2 == player && round.Grid.Position3 == player {
		winner = player
	} else if round.Grid.Position4 == player && round.Grid.Position5 == player && round.Grid.Position6 == player {
		winner = player
	} else if round.Grid.Position7 == player && round.Grid.Position8 == player && round.Grid.Position9 == player {
		winner = player
	} else if round.Grid.Position1 == player && round.Grid.Position5 == player && round.Grid.Position9 == player {
		winner = player
	} else if round.Grid.Position3 == player && round.Grid.Position5 == player && round.Grid.Position7 == player {
		winner = player
	} else if round.Grid.Position1 == player && round.Grid.Position4 == player && round.Grid.Position7 == player {
		winner = player
	} else if round.Grid.Position2 == player && round.Grid.Position5 == player && round.Grid.Position8 == player {
		winner = player
	} else if round.Grid.Position3 == player && round.Grid.Position6 == player && round.Grid.Position9 == player {
		winner = player
	}
	return winner
}

func messageSelect() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(3)
	message := playingMessages[index]
	return message
}
