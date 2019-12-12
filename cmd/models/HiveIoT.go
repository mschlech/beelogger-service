package models

import (
)

type HiveIoT struct {
	Id int `json:"id"`
	battery string `json:"battery"`
	batteryConsumption string `json:"batteryConsumption"`
	firmwareVersion int `json:"firmwareVersion"`
}
