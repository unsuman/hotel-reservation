package db

var (
	DBuri      string
	DBname     string
	TestDBname string
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
