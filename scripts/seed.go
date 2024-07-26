package main

import (
	"context"
	"fmt"
	"log"

	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBname)
	roomStore := db.NewMongoRoomStore(client, hotelStore, db.DBname)

	hotel := types.Hotel{
		Name:     "Mayfair",
		Location: "Puri",
		Rooms:    []primitive.ObjectID{},
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

	if err := client.Database(db.DBname).Drop(ctx); err != nil {
		log.Fatal(err)
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

		fmt.Println(insertedRoom)

	}

	fmt.Println(insertedHotel)
}
