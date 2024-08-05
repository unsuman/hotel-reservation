package db

import (
	"context"
	"log"

	"github.com/unsuman/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingStore interface {
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
	GetBookings(context.Context, bson.M) ([]*types.Booking, error)
	GetBookingByID(context.Context, string) (*types.Booking, error)
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client, dbname string) *MongoBookingStore {
	return &MongoBookingStore{
		client: client,
		coll:   client.Database(dbname).Collection("bookings"),
	}
}

func (b *MongoBookingStore) GetBookingByID(ctx context.Context, id string) (*types.Booking, error) {
	var booking types.Booking
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := b.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&booking); err != nil {
		return nil, err
	}

	return &booking, nil
}

func (b *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error) {
	cur, err := b.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var bookings []*types.Booking
	if err = cur.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (b *MongoBookingStore) InsertBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	res, err := b.coll.InsertOne(ctx, booking)
	if err != nil {
		log.Fatal(err)
	}

	booking.ID = res.InsertedID.(primitive.ObjectID)

	return booking, nil
}
