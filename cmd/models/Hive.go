package models

import (
)

type Hive struct {
	Id int `json:"id"`
	temperature string `json:"temperature"`
	humidity int `json:"humidity"`
	lux int `json:"lux"`
	weight int `json:"weight"`
	location string `json:"location"`
	gas string `json:"gas"`
}



