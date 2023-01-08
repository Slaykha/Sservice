package models

import "time"

type Poll struct {
	ID           string    `json:"id" bson:"id"`
	UserId       string    `json:"userId" bson:"userId"`
	PollQuestion string    `json:"pollQuestion" bson:"pollQuestion"`
	PollOptions  []Options `json:"options" bson:"options"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
}

type PollDTO struct {
	PollQuestion string    `json:"pollQuestion" bson:"pollQuestion"`
	PollOptions  []Options `json:"options" bson:"options"`
}

type Options struct {
	OptionText  string `json:"optionText" bson:"optionText"`
	OptionVotes int    `json:"optionVotes" bson:"optionVotes"`
}

type User struct {
	ID           string    `json:"id" bson:"id"`
	Name         string    `json:"name" bson:"name"`
	Email        string    `json:"email" bson:"email"`
	Password     []byte    `json:"password" bson:"password"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	UserVotes    []Votes   `json:"userVotes" bson:"userVotes"`
	ProfilePhoto string    `json:"profilePhoto" bson:"profilePhoto"`
}

type UserRegisterDTO struct {
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"password"`
	ProfilePhoto string `json:"profilePhoto" bson:"profilePhoto"`
}

type UserLoginDTO struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Votes struct {
	PollId string `json:"pollId" bson:"pollId"`
	Vote   string `json:"vote" bson:"vote"`
}
