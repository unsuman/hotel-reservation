package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/types"
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
	user := c.Context().UserValue("user").(*types.User)

	booking, err := b.store.BookingStore.GetBookingByID(c.Context(), id)
	if err != nil {
		return err
	}

	if user.ID != booking.UserID {
		c.SendStatus(fiber.StatusUnauthorized)
		return fmt.Errorf("not authorized")
	}

	return c.JSON(booking)
}

func (b *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	user := c.Context().UserValue("user").(*types.User)

	var filter bson.M
	if !user.IsAdmin {
		filter = bson.M{
			"userID": user.ID,
		}
	} else {
		filter = bson.M{}
	}
	booking, err := b.store.BookingStore.GetBookings(c.Context(), filter)
	if err != nil {
		return err
	}

	if len(booking) == 0 {
		return fmt.Errorf("no bookings")
	}

	return c.JSON(booking)
}
