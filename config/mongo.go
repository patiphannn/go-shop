package config

import (
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {
	host := viper.Get("MONGO_HOST").(string)
	dbName := viper.Get("MONGO_DB_NAME").(string)

	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)
	return db, nil
}
