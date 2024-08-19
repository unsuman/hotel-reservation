package db

import "os"

var (
	DBuri      = os.Getenv("MONGO_DB_URL")
	DBname     = os.Getenv("MONGO_DB_NAME")
	TestDBname = os.Getenv("MONGO_TEST_DB_NAME")
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
