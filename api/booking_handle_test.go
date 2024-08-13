package api

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db/fixtures"
	"github.com/unsuman/hotel-reservation.git/types"
)

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)

	defer db.teardown(t)

	var (
		user       = fixtures.AddUser(db.store, "foo", "bar", false)
		user_admin = fixtures.AddUser(db.store, "admin", "admin", true)

		hotel1 = fixtures.AddHotel(db.store, "Mayfair", "Mumbai", 4, nil)
		room1  = fixtures.AddRoom(db.store, "large", true, 5000.0, hotel1.ID)
		room2  = fixtures.AddRoom(db.store, "medium", false, 3000.0, hotel1.ID)

		booking_user       = fixtures.AddBooking(db.store, user.ID, room1.HotelID, time.Now(), time.Now().AddDate(0, 0, 4))
		booking_user_admin = fixtures.AddBooking(db.store, user_admin.ID, room2.HotelID, time.Now(), time.Now().AddDate(0, 0, 3))

		app = fiber.New()

		bookingHandler = NewBookingHandler(db.store)
	)
	userRoute := app.Group("/", JWTAuthentication(db.store.UserStore))

	admin := userRoute.Group("/admin", AdminAuthorization())
	userRoute.Get("/bookings", bookingHandler.HandleGetBookings)
	admin.Get("/bookings", bookingHandler.HandleGetBookings)

	req := httptest.NewRequest("GET", "/admin/bookings", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user_admin))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	var bookings []types.Booking
	_, _ = booking_user, booking_user_admin

	json.NewDecoder(resp.Body).Decode(&bookings)
	if len(bookings) != 2 {
		t.Fatalf("expected number of bookings for admin to be 2 but got %v", len(bookings))
	}

	req = httptest.NewRequest("GET", "/bookings", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))

	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	json.NewDecoder(resp.Body).Decode(&bookings)

	if len(bookings) != 1 {
		t.Fatalf("expected number of bookings for user to be 1 but got %v", len(bookings))
	}

}
