package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("-- JWT Authing")

	token, ok := c.GetReqHeaders()["X-Api-Token"]
	fmt.Println("-- (1)")
	if !ok {
		return fmt.Errorf("unauthorized")
	}

	if err := parseToken(token[0]); err != nil {
		return fmt.Errorf("Unauthorized")
	}
	fmt.Println("-- (2)")
	fmt.Println("Token: ", token)
	return nil
}

func parseToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("Unauthorized")
		}
		fmt.Println("-- (3)")

		secret := os.Getenv("JWT_SECRET")

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	fmt.Println(tokenStr)
	if err != nil {
		fmt.Println(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims)
	}
	return fmt.Errorf("Unauthorized")
}
