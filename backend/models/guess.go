package models

type Guess struct {
	Guess []string `json:"Guess"`
}

type GuessResults struct {
	Guess   []string `json:"Guess"`
	Present []bool   `json:"Present"`
	Correct []bool   `json:"Correct"`
}
