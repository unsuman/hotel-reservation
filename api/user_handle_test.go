package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	store *db.Store
}

func setup() *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		store: &db.Store{
			UserStore: db.NewMongoUserStore(client, db.UserTestDBname),
		},
	}
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.store.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestUserCRUD(t *testing.T) {

	tdb := setup()

	app := fiber.New()
	userHandler := NewUserHandler(tdb.store)

	defer setup().teardown(t)

	var (
		user   types.User
		userID string
	)
	t.Run("TestPostUser", func(t *testing.T) {

		app.Post("/", userHandler.HandlePostUser)

		params := types.CreateUsersParams{
			FirstName: "Ansuman",
			LastName:  "At the water cooler",
			Email:     "water@water.com",
			Pass:      "aadnmipkmm2434",
		}

		b, _ := json.Marshal(params)

		req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
		req.Header.Add("Content-type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}

		json.NewDecoder(resp.Body).Decode(&user)

		if len(user.ID) == 0 {
			t.Errorf("expected an user id")
		}
		if user.FirstName != params.FirstName {
			t.Errorf("expected first name %s but got %s", params.FirstName, user.FirstName)
		}

		if user.LastName != params.LastName {
			t.Errorf("expected last name %s but got %s", params.LastName, user.LastName)
		}
		if user.Email != params.Email {
			t.Errorf("expected email %s but got %s", params.Email, user.Email)
		}
	})

	t.Run("TestGetUser", func(t *testing.T) {
		app.Get("/:id", userHandler.HandleGetUser)

		userID = "/" + user.ID.Hex()

		req := httptest.NewRequest("GET", userID, nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}

		var singleUser types.User

		json.NewDecoder(resp.Body).Decode(&singleUser)

		if len(user.ID) == 0 {
			t.Errorf("expected an user id")
		}
		if user.FirstName != singleUser.FirstName {
			t.Errorf("expected first name %s but got %s", singleUser.FirstName, user.FirstName)
		}

		if user.LastName != singleUser.LastName {
			t.Errorf("expected last name %s but got %s", singleUser.LastName, user.LastName)
		}
		if user.Email != singleUser.Email {
			t.Errorf("expected email %s but got %s", singleUser.Email, user.Email)
		}

	})

	type Status struct {
		Updated string `json:"updated"`
	}

	t.Run("TestUpdateUser", func(t *testing.T) {
		app.Put("/:id", userHandler.HandleUpdateUser)

		userUpdate := &types.UpdateUserParams{
			FirstName: "Aditya",
			LastName:  "At the beach",
		}

		b, _ := json.Marshal(userUpdate)

		req := httptest.NewRequest("PUT", userID, bytes.NewReader(b))
		req.Header.Add("Content-type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}

		var stat Status

		json.NewDecoder(resp.Body).Decode(&stat)

		if stat.Updated != user.ID.Hex() {
			t.Errorf("Error updating %s", user.ID.Hex())
		}
	})

}
