package api

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type roomHandler struct {
	store db.Store
}

func NewRoomHandler(store *db.Store) *roomHandler {
	return &roomHandler{
		store: *store,
	}
}

type BookingRes struct {
	NumPersons int       `json:"numPersons,omitempty"`
	FromDate   time.Time `json:"fromDate,omitempty"`
	TillDate   time.Time `json:"tillDate,omitempty"`
}

func (r *roomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var res BookingRes
	if err := c.BodyParser(&res); err != nil {
		return err
	}
	err := res.validate()
	if err != nil {
		return err
	}

	roomID := c.Params("id")
	roomOID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return err
	}

	user := c.Context().Value("user").(*types.User)

	booking := types.Booking{
		RoomID:     roomOID,
		UserID:     user.ID,
		FromDate:   res.FromDate,
		TillDate:   res.TillDate,
		NumPersons: res.NumPersons,
	}

	ok, err := r.IsRoomAvailable(c, &booking)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("no rooms available")
	}

	booked, err := r.store.BookingStore.InsertBooking(c.Context(), &booking)
	if err != nil {
		return err
	}

	fmt.Println("booked:", booked)

	return c.JSON(booked)

}

func (b BookingRes) validate() error {
	now := time.Now()
	if now.After(b.FromDate) || now.After(b.TillDate) {
		return fmt.Errorf("cannot book a room in the past")
	}
	return nil
}

func (r *roomHandler) IsRoomAvailable(c *fiber.Ctx, booking *types.Booking) (bool, error) {
	filter := bson.M{
		"roomID": booking.RoomID,
		"fromDate": bson.M{
			"$gte": booking.FromDate,
		},
		"tillDate": bson.M{
			"$lte": booking.TillDate,
		},
	}
	rooms, err := r.store.BookingStore.GetBookings(c.Context(), filter)
	if err != nil {
		log.Fatal(err)
		return true, err
	}

	if len(rooms) != 0 {
		return false, nil
	}

	return true, nil

}
