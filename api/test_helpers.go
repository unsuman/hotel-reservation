package api

import (
	"context"
	"log"
	"testing"

	"github.com/unsuman/hotel-reservation.git/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	client *mongo.Client
	store  *db.Store
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.TestDBname)

	return &testdb{
		store: &db.Store{
			UserStore:    db.NewMongoUserStore(client, db.TestDBname),
			RoomStore:    db.NewMongoRoomStore(client, hotelStore, db.TestDBname),
			HotelStore:   db.NewMongoHotelStore(client, db.TestDBname),
			BookingStore: db.NewMongoBookingStore(client, db.TestDBname),
		},
		client: client,
	}
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.store.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
	tdb.client.Database(db.TestDBname).Drop(context.Background())
}
