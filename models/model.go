package models

import "time"

type Poll struct {
	ID           string    `json:"id"`
	UserId       string    `json:"userId"`
	PollQuestion string    `json:"pollQuestion"`
	PollOptions  []Options `json:"options"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
}

type Options struct {
	OptionText  string `json:"optionText"`
	OptionVotes int    `json:"optionVotes"`
}

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     []byte    `json:"password"`
	CreatedAt    time.Time `json:"createdAt"`
	UserVotes    []Votes   `json:"userVotes"`
	ProfilePhoto string    `json:"profilePhoto"`
}

type Votes struct {
	PollId string `json:"pollId"`
	Vote   string `json:"vote"`
}
