package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/unsuman/hotel-reservation.git/api"
	"github.com/unsuman/hotel-reservation.git/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	var (
		app              = fiber.New(config)
		apiv1WithoutAuth = app.Group("/api")

		hotelStore   = db.NewMongoHotelStore(client, db.DBname)
		roomStore    = db.NewMongoRoomStore(client, hotelStore, db.DBname)
		userStore    = db.NewMongoUserStore(client, db.DBname)
		bookingStore = db.NewMongoBookingStore(client, db.DBname)

		store = &db.Store{
			HotelStore:   hotelStore,
			RoomStore:    roomStore,
			UserStore:    userStore,
			BookingStore: bookingStore,
		}

		apiv1         = app.Group("/api/v1", api.JWTAuthentication(userStore))
		admin         = apiv1.Group("/admin", api.AdminAuthorization())
		userHandler   = api.NewUserHandler(store)
		hotelHandler  = api.NewHotelHandler(store)
		authHandler   = api.NewAuthHandler(userStore)
		roomHandler   = api.NewRoomHandler(store)
		bookingHander = api.NewBookingHandler(store)
	)

	//auth
	apiv1WithoutAuth.Post("/auth", authHandler.HandleAuthentication)

	// user handlers
	apiv1.Put("/user/:id", userHandler.HandleUpdateUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Post("/user", userHandler.HandlePostUser)

	// hotel handlers
	apiv1.Get("/hotels", hotelHandler.GetHotels)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.GetRooms)
	apiv1.Get("/hotel/:id", hotelHandler.GetHotel)

	//booking handlers
	apiv1.Post("/room/:id/book", roomHandler.HandleBookRoom)
	apiv1.Get("/booking/:id", bookingHander.HandleGetBooking)
	apiv1.Get("/bookings", bookingHander.HandleGetBookings)
	apiv1.Get("/booking/:id/cancel", bookingHander.HandleCancelBooking)
	admin.Get("/bookings", bookingHander.HandleGetBookings)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
