package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
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
