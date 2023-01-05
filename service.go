package main

import (
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func createID() (id string) {
	id = uuid.New().String()

	id = strings.ReplaceAll(id, "-", "")

	id = id[0:8]

	return
}
