package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

	user, err := h.userStore.GetUserByEmail(c.Context(), authParams.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("invalid credentials")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPass), []byte(authParams.Password))
	if err != nil {
		return fmt.Errorf("invalid credentials")
	}

	fmt.Println("Authenticated the user -> ", user)

	return nil
}
