package models

type Poll struct {
	ID           int       `json:"id"`
	PollQuestion string    `json:"pollQuestion"`
	PollOptions  []Options `json:"options"`
}

type Options struct {
	OptionText string `json:"optionText"`
	Responses  int    `json:"responses"`
}
