package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/types"
)

func AdminAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Context().UserValue("user").(*types.User)
		if !user.IsAdmin {
			return fmt.Errorf("unauthorized request")
		}

		return c.Next()
	}
}
