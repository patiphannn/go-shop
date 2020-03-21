package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Shop is field shop
type Shop struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `bson:"name" form:"name" json:"name" xml:"name" binding:"required"`
	Detail      string        `bson:"detail" form:"detail" json:"detail" xml:"detail" binding:"required"`
	CreatedTime time.Time     `json:"created_time" bson:"created_time"`
	UpdatedTime time.Time     `json:"updated_time" bson:"updated_time"`
}

// Shops is multiple shop
type Shops []Shop
