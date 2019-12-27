package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HiveIoT struct {
	ID primitive.ObjectID `json:"_id"`
	battery string `json:"battery"`
	batteryConsumption string `json:"batteryConsumption"`
	firmwareVersion int `json:"firmwareVersion"`
}
