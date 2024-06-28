package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "unsuman",
		LastName:  "sahoo",
	}
	return c.JSON(u)
}
