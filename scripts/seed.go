package main

import (
	"context"
	"log"

	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx        = context.Background()
	client     *mongo.Client
	hotelStore db.HotelStore
	roomStore  db.RoomStore
	userStore  db.UserStore
)

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBname).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore = db.NewMongoHotelStore(client, db.DBname)
	roomStore = db.NewMongoRoomStore(client, hotelStore, db.DBname)
	userStore = db.NewMongoUserStore(client, db.DBname)
}

func seedUser(fname string, lname string, email string) {
	user, err := types.NewUserFromParams(types.CreateUsersParams{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Pass:      "verysecurepass",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = userStore.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
}

func seedHotel(hotelName string, location string, rating int) {
	hotel := types.Hotel{
		Name:     hotelName,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []types.Room{
		{
			Type:      types.DeluxeRoom,
			BasePrice: 4343.6,
		},
		{
			Type:      types.DoubleRoom,
			BasePrice: 1212.12,
		},
		{
			Type:      types.SeaSideRoom,
			BasePrice: 2323.12,
		},
		{
			Type:      types.SingleRoom,
			BasePrice: 1000.0,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID

		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		_ = insertedRoom
	}
}

func main() {
	seedUser("James", "At the water cooler", "foo@bar.com")
	seedHotel("Mayfair", "Puri", 4)
	seedHotel("Pal heights", "Bhubaneswar", 3)
	seedHotel("Taj", "Mumbai", 5)
}
