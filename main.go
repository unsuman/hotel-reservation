package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/api"
	"github.com/unsuman/hotel-reservation.git/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New(config)

	apiv1 := app.Group("/api/v1")

	hotelStore := db.NewMongoHotelStore(client, db.DBname)
	roomStore := db.NewMongoRoomStore(client, hotelStore, db.DBname)

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, db.DBname))
	hotelHandler := api.NewHotelHandler(hotelStore, roomStore)

	apiv1.Put("/user/:id", userHandler.HandleUpdateUser)

	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)

	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	apiv1.Post("/user", userHandler.HandlePostUser)

	apiv1.Get("/hotels", hotelHandler.GetHotels)
	app.Listen(*listenAddr)
}
