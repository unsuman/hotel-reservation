package api

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/unsuman/hotel-reservation.git/db"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok {
			return ErrUnAuthorized()
		}

		claims, err := validateToken(token[0])
		if err != nil {
			return ErrUnAuthorized()
		}
		expires, _ := claims["exp"].(float64)
		if time.Now().Unix() > int64(expires) {
			return fmt.Errorf("token has expired")
		}

		userID := claims["id"].(string)

		user, err := userStore.GetUserByID(context.TODO(), userID)
		if err != nil {
			ErrInvalidID()
		}

		c.Context().SetUserValue("user", user)

		return c.Next()
	}

}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, ErrUnAuthorized()
		}

		secret := os.Getenv("JWT_SECRET")

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrUnAuthorized()
	}
	return claims, nil
}
