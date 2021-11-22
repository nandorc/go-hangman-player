package model

type Request struct {
	WordToGuess string `json:"palabra"`
	Tries       string `json:"intentos"`
}
