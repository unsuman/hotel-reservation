package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
	Rating   int                  `json:"rating" bson:"rating"`
}

type Room struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Size    string             `bson:"size" json:"size"`
	Seaside bool               `bson:"seaside" json:"seaside"`
	Price   float64            `json:"price" bson:"price"`
	HotelID primitive.ObjectID `json:"hotelID" bson:"hotelID"`
}
