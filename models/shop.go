package model

import (
	"time"
)

type Shop struct {
	name      string    `bson:"name"`
	detail    string    `bson:"detail"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type Shops []Shop
