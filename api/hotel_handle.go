package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type hotelHandler struct {
	store db.Store
}

func NewHotelHandler(store *db.Store) *hotelHandler {
	return &hotelHandler{
		store: *store,
	}
}

func (s *hotelHandler) GetHotel(c *fiber.Ctx) error {
	id := c.Params("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}

	hotel, err := s.store.HotelStore.GetHotelByID(c.Context(), oid)
	if err != nil {
		return err
	}

	return c.JSON(hotel)

}
func (s *hotelHandler) GetHotels(c *fiber.Ctx) error {
	hotels, err := s.store.HotelStore.GetHotels(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}

func (s *hotelHandler) GetRooms(c *fiber.Ctx) error {
	id := c.Params("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"hotelID": oid}
	hotels, err := s.store.RoomStore.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}
