package db

const (
	DBuri      = "mongodb://localhost:27017"
	DBname     = "hotel-reservation"
	TestDBname = "hotel-reservation-test"
)

type Store struct {
	UserStore    UserStore
	HotelStore   HotelStore
	RoomStore    RoomStore
	BookingStore BookingStore
}
