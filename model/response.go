package model

type Response struct {
	Character string `json:"letra"`
}

func (response *Response) Build(previous string, next string, exclusions string) {
	var suggestion string = GetSuggestion(exclusions)
	response.Character = suggestion
}
