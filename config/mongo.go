package config

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

// GetMongoDB is setup mongodb
func GetMongoDB() (*mgo.Database, error) {
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")

	if host == "" {
		host = "localhost:27017"
	}

	if dbName == "" {
		dbName = "go-shop"
	}

	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)
	return db, nil
}
