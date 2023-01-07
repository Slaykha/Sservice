package main

import (
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

	err := c.BodyParser(&UserDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	user, err := a.service.LoginUser(UserDTO)

	switch err {
	case nil:
		c.JSON(user)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}
