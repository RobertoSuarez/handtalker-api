package main

import (
	"log"

	"github.com/RobertoSuarez/vinculacion_api_graph/db"
	"github.com/RobertoSuarez/vinculacion_api_graph/handlers"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {

	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	api := app.Group("/api")

	userHandlers := handlers.NewUserHandler()

	users := api.Group("/users")
	users.Get("/", userHandlers.HandleGetUsers)
	users.Post("/", userHandlers.HandlePostUser)

	log.Fatal(app.Listen(":4000"))
}
