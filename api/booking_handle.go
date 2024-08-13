package api

import (
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

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.BookingStore.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	if err := h.store.BookingStore.UpdateBooking(c.Context(), c.Params("id"), bson.M{"canceled": true}); err != nil {
		return err
	}
	return c.JSON("updated")
}

func (b *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	user := c.Context().UserValue("user").(*types.User)

	booking, err := b.store.BookingStore.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("bookings")
	}

	if user.ID != booking.UserID {
		c.SendStatus(fiber.StatusUnauthorized)
		return ErrUnAuthorized()
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
		return ErrNotResourceNotFound("bookings")
	}

	return c.JSON(booking)
}
