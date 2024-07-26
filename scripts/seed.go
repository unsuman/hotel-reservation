package main

import (
	"context"
	"fmt"
	"log"

	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/types"
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
	roomStore := db.NewMongoRoomStore(client, db.DBname)

	hotel := types.Hotel{
		Name:     "Mayfair",
		Location: "Puri",
	}

	room := types.Room{
		Type:      types.DeluxeRoom,
		BasePrice: 4343.6,
	}

	_ = room

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	room.HotelID = insertedHotel.ID

	insertedRoom, err := roomStore.InsertRoom(ctx, &room)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertedRoom)

	fmt.Println(insertedHotel)
}
