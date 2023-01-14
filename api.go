package main

import (
	"github.com/Slaykha/Poll-App-Service/helpers"
	"github.com/Slaykha/Poll-App-Service/models"
	"github.com/gofiber/fiber"
)

type Api struct {
	service *Service
}

func NewAPI(service *Service) *Api {
	return &Api{
		service: service,
	}
}

func (a *Api) HandleUserCreate(c *fiber.Ctx) {
	var UserDTO models.UserRegisterDTO

	err := c.BodyParser(&UserDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	user, err := a.service.CreateUser(UserDTO)

	switch err {
	case nil:
		c.JSON(user)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (a *Api) HandleUserLogin(c *fiber.Ctx) {
	var UserDTO models.UserLoginDTO
	var token string

	err := c.BodyParser(&UserDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	user, err := a.service.LoginUser(UserDTO)
	if err == nil {
		token, _ = helpers.CreateToken(user.ID)
	}

	switch err {
	case nil:
		c.JSON(token)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (a *Api) HandlePollCreate(c *fiber.Ctx) {
	UserID := c.Params("userId")

	var PollDTO models.PollDTO

	err := c.BodyParser(&PollDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	user, err := a.service.CreatePoll(PollDTO, UserID)

	switch err {
	case nil:
		c.JSON(user)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (a *Api) HandlePollsGet(c *fiber.Ctx) {
	polls, err := a.service.GetPolls()

	switch err {
	case nil:
		c.JSON(polls)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

}
