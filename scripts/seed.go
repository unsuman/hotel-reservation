package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/unsuman/hotel-reservation.git/api"
	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/db/fixtures"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBname).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBname)
	store := &db.Store{
		UserStore:    db.NewMongoUserStore(client, db.DBname),
		BookingStore: db.NewMongoBookingStore(client, db.DBname),
		RoomStore:    db.NewMongoRoomStore(client, hotelStore, db.DBname),
		HotelStore:   hotelStore,
	}

	user := fixtures.AddUser(store, "unsuman", "bar", false)
	fmt.Println("unsuman ->", api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("admin ->", api.CreateTokenFromUser(admin))
	hotel := fixtures.AddHotel(store, "Taj", "Mumbai", 5, nil)
	room := fixtures.AddRoom(store, "large", true, 88.44, hotel.ID)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))
	fmt.Println("booking ->", booking.ID)
}
