package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go/log"
	"github.com/unsuman/hotel-reservation.git/types"
)

func TestUserAuthFailureWithWrongPass(t *testing.T) {
	tdb, _ := seedTestDB()

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.store.UserStore)

	defer tdb.teardown(t)

	app.Post("/auth", authHandler.HandleAuthentication)

	params := AuthParams{
		Email:    "foo@bar.com",
		Password: "verysecurepassnotcorrect",
	}

	b, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", "/auth", bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("status code: ", fiber.StatusBadRequest, resp.StatusCode)

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected http status of 400 but got %d", resp.StatusCode)
	}
}

func seedTestDB() (*testdb, *types.User) {
	tdb := setup()

	user, _ := types.NewUserFromParams(types.CreateUsersParams{
		FirstName: "James",
		LastName:  "At the water cooler",
		Email:     "foo@bar.com",
		Pass:      "verysecurepass",
	})
	insertedUser, err := tdb.store.UserStore.InsertUser(context.TODO(), user)
	if err != nil {
		log.Error(err)
	}
	return tdb, insertedUser
}

func TestUserAuthSuccess(t *testing.T) {

	tdb, insertedUser := seedTestDB()

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.store.UserStore)

	defer tdb.teardown(t)

	app.Post("/auth", authHandler.HandleAuthentication)

	params := AuthParams{
		Email:    "foo@bar.com",
		Password: "verysecurepass",
	}

	b, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", "/auth", bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode == fiber.StatusBadRequest {
		t.Fatalf("expected status code of 200 but got %d", resp.StatusCode)
	}

	var authResp Resp
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Fatal(err)
	}

	if authResp.Token == "" {
		t.Fatalf("expected the JWT token in the response")
	}

	if insertedUser.ID != authResp.User.ID {
		t.Fatalf("user not found")
	}

}
