package main

import (
	"fmt"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

type Config struct {
	AppPort         int
	Host            string
	DBReplicaSetUrl string
}

var config Config

func main() {
	fmt.Println("Service Starting...")
	setConfig()

	repository := NewRepository(config.DBReplicaSetUrl)
	service := NewService(repository)
	API := NewAPI(service)

	app := SetupApp(API)

	fmt.Println("Poll App service started at ", config.AppPort, "  ...")

	app.Post("/user/register", API.HandleUserCreate)
	app.Post("/user/login", API.HandleUserLogin)

	app.Post("/poll/create/:userId", API.HandlePollCreate)
	app.Get("/polls", API.HandlePollsGet)
	/* app.Get("/poll/list/:UserId")
	app.Put("/poll/:pollId/vote/:userID")
	app.Delete("/poll/:pollId/delete") */

	app.Get("/status", func(c *fiber.Ctx) {
		c.Status(fiber.StatusOK)
	})

	app.Listen(config.AppPort)
}

func SetupApp(API *Api) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{"Origin, Content-Type, Accept"},
	}))

	return app
}

func setConfig() {
	config = Config{
		AppPort:         12345,
		Host:            "192.168.1.101:12345",
		DBReplicaSetUrl: "mongodb+srv://admin:HkJpLyv1MclTvMIc@spendingtraacker.ybzvy6n.mongodb.net/?retryWrites=true&w=majority",
	}
}
