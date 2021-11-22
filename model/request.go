package model

import "strings"

type Request struct {
	WordToGuess string `json:"palabra"`
	Tries       string `json:"intentos"`
}

func (request *Request) GetFirstUnkown() (bool, string, string) {
	index := strings.Index(request.WordToGuess, "?")
	if index < 0 {
		return false, "", ""
	} else if index == 0 {
		if len(request.WordToGuess) <= 1 {
			return true, "", ""
		} else {
			return true, "", string(request.WordToGuess[1])
		}
	} else if index == len(request.WordToGuess)-1 {
		return true, string(request.WordToGuess[index-1]), ""
	} else {
		return true, string(request.WordToGuess[index-1]), string(request.WordToGuess[index+1])
	}
}
