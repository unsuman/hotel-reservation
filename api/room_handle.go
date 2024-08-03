package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/types"
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

	roomID := c.Params("id")
	roomOID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return err
	}

	user := c.Context().Value("user").(*types.User)

	fmt.Printf("%+v \n", &types.Booking{
		RoomID:     roomOID,
		UserID:     user.ID,
		FromDate:   res.FromDate,
		TillDate:   res.TillDate,
		NumPersons: res.NumPersons,
	})

	return nil

}
