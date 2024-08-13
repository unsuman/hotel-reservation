package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db/fixtures"
)

func TestUserAuthFailureWithWrongPass(t *testing.T) {
	tdb := setup(t)

	_ = fixtures.AddUser(tdb.store, "James", "Foo", false)
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
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

	if resp.StatusCode == http.StatusBadRequest {
		t.Fatalf("expected http status of 400 but got %d", resp.StatusCode)
	}
}

func TestUserAuthSuccess(t *testing.T) {

	tdb := setup(t)
	user := fixtures.AddUser(tdb.store, "James", "Foo", false)

	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
	authHandler := NewAuthHandler(tdb.store.UserStore)

	defer tdb.teardown(t)

	app.Post("/auth", authHandler.HandleAuthentication)

	params := AuthParams{
		Email:    "James@Foo.com",
		Password: "James_Foo",
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

	if user.ID != authResp.User.ID {
		t.Fatalf("user not found")
	}

}
