package db

const (
	DBuri      = "mongodb://localhost:27017"
	DBname     = "hotel-reservation"
	TestDBname = "hotel-reservation-test"
)

type Pagination struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}

type Store struct {
	UserStore    UserStore
	HotelStore   HotelStore
	RoomStore    RoomStore
	BookingStore BookingStore
}
