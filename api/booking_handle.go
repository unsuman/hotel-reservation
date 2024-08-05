package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: *store,
	}
}

func (b *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := b.store.BookingStore.GetBookingByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(booking)
}

func (b *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	booking, err := b.store.BookingStore.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return err
	}

	return c.JSON(booking)
}
