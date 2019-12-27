package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "gopkg.in/mgo.v2/bson"
)
type Hive struct {
	ID primitive.ObjectID `bson:"_id json:"_id"`
	temperature string `bson:"temperature json:"temperature"`
	humidity int `bson:"humidity json:"humidity"`
	lux int `bson:"lux json:"lux"`
	weight int `bson:"weight json:"weight"`
	location string `bson:"location json:"location"`
}


