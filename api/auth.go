package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
)

type UserAuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *UserAuthHandler {
	return &UserAuthHandler{
		userStore: userStore,
	}
}

type AuthParams struct {
	Email    string
	Password string
}

func (h *UserAuthHandler) HandleAuthentication(c *fiber.Ctx) error {
	var authParams AuthParams
	if err := c.BodyParser(&authParams); err != nil {
		return err
	}

	fmt.Println(authParams)

	return nil
}
