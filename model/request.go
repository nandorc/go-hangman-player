package model

type Request struct {
	WordToGuess string `json:"palabra"`
	TriesCount  string `json:"intentos"`
}
