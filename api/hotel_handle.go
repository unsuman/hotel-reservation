package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type hotelHandler struct {
	db.HotelStore
	db.RoomStore
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *hotelHandler {
	return &hotelHandler{
		HotelStore: hs,
		RoomStore:  rs,
	}
}

func (s *hotelHandler) GetHotels(c *fiber.Ctx) error {
	hotels, err := s.HotelStore.GetHotels(c.Context())
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
	hotels, err := s.RoomStore.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}
