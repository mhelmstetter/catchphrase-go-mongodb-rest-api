package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Catchphrase struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RevenueType   string             `json:"revenueType,omitempty" bson:"revenueType,omitempty"`
	AdvertisingID string             `json:"advertisingID,omitempty" bson:"advertisingID,omitempty"`
	DeviceID      string             `json:"deviceID,omitempty" bson:"deviceID,omitempty"`
}
