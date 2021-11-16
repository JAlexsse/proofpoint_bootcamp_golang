package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/game"
	"net/http"
	"strconv"
)

var roundGrid = game.Grid{
	Position1: 0,
	Position2: 0,
	Position3: 0,
	Position4: 0,
	Position5: 0,
	Position6: 0,
	Position7: 0,
	Position8: 0,
	Position9: 0,
}

var round = game.Round{
	Grid:    roundGrid,
	Turn:    0,
	Winner:  0,
	IsDraw:  false,
	Message: "",
}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/again", restartGame)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Starting web wever on port 8080")
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("playerChoice"))
	game.PlayRound(playerChoice, &round)
	out, err := json.MarshalIndent(round, "", "  ")
	if err != nil {
		log.Println("err")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)

	if err != nil {
		log.Println("err")
		return
	}
	err = t.Execute(w, nil)

	if err != nil {
		log.Println("err")
		return
	}
}

func restartGame(w http.ResponseWriter, r *http.Request) {
	round.Grid.Position1 = 0
	round.Grid.Position2 = 0
	round.Grid.Position3 = 0
	round.Grid.Position4 = 0
	round.Grid.Position5 = 0
	round.Grid.Position6 = 0
	round.Grid.Position7 = 0
	round.Grid.Position8 = 0
	round.Grid.Position9 = 0
	round.Message = ""
	round.Turn = 0
	round.IsDraw = false
	round.Winner = 0
	renderTemplate(w, "index.html")
}
