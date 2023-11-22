package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
)

func checkGuess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	reqBody, _ := io.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))
	answer := []string{"c", "a", "t"}

	var g models.Guess
	err := json.Unmarshal(reqBody, &g)
	if err != nil {
		fmt.Println("Unmarshal had an error", string(reqBody))
		return
	}

	if len(answer) != len(g.Guess) {
		fmt.Println("the given guess is the wrong length", g.Guess)
		return
	}

	var result models.GuessResults
	result.Guess = g.Guess
	result.Present = make([]bool, len(result.Guess))
	result.Correct = make([]bool, len(result.Guess))

	for i, letter := range result.Guess {
		if letter == answer[i] {
			result.Correct[i] = true
		}

		for j := 0; j < len(answer); j++ {
			if letter == answer[j] {
				result.Present[i] = true
				break
			}
		}
	}

	_ = json.NewEncoder(w).Encode(result)
}
