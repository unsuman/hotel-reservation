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

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type QueryParams struct {
	db.Pagination
	Rating int `json:"rating"`
}

func (s *hotelHandler) GetHotels(c *fiber.Ctx) error {
	var params QueryParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}

	filter := bson.M{}
	if params.Rating != 0 {
		filter = bson.M{"rating": params.Rating}
	}
	hotels, err := s.store.HotelStore.GetHotels(c.Context(), filter, &params.Pagination)
	if err != nil {
		return err
	}
	resp := ResourceResp{
		Data:    hotels,
		Results: len(*hotels),
		Page:    int(params.Page),
	}
	return c.JSON(resp)
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
