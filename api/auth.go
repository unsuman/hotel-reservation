package api

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/mongo"
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

type Resp struct {
	User  *types.User
	Token string
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

	ok := types.IsPasswordValid(user.EncryptedPass, authParams.Password)
	if !ok {
		return fmt.Errorf("invalid credentials")
	}

	fmt.Println("Authenticated the user -> ", user)

	res := &Resp{
		User:  user,
		Token: h.createTokenFromUser(user),
	}
	return c.JSON(res)
}

func (h *UserAuthHandler) createTokenFromUser(user *types.User) string {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.FirstName + user.LastName,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_Secret")
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	return t
}
