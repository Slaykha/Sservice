package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Slaykha/Poll-App-Service/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateUser(UserDTO models.UserRegisterDTO) (*models.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(UserDTO.Password), 8)

	emptyVotes := []models.Votes{}

	user := models.User{
		ID:           generateID(),
		Name:         UserDTO.Name,
		Email:        UserDTO.Email,
		Password:     password,
		CreatedAt:    time.Now().UTC(),
		UserVotes:    emptyVotes,
		ProfilePhoto: UserDTO.ProfilePhoto,
	}

	err := s.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) LoginUser(userDTO models.UserLoginDTO) (*models.User, error) {
	user, err := s.repository.FindUser(userDTO.Email)
	if err != nil {
		fmt.Println("1", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.Password))
	if err != nil {
		fmt.Println("2", err)
		return nil, err
	}

	return user, nil
}

func generateID() (id string) {
	id = uuid.New().String()

	id = strings.ReplaceAll(id, "-", "")

	id = id[0:8]

	return
}
