package handlers

import (
	"net/http"

	"github.com/RobertoSuarez/vinculacion_api_graph/data"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {

	userRequest := &data.CreateUserRequest{}

	if err := c.BodyParser(userRequest); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := userRequest.Validate(); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// creamos el usuarios
	user, err := data.CreateUser(userRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	// enviamos el usuario creado
	return c.Status(200).JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	users, err := data.GetUsers()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(users)
}
