package main

import (
	"log"

	mongo "github.com/polnoy/go-shop/config"
	routes "github.com/polnoy/go-shop/routes"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	connectionDB, err := mongo.GetMongoDB()
	if err != nil {
		log.Panic("Can no connect Database", err.Error())
	}

	r := routes.SetupRouter(connectionDB)
	r.Run(":8080")
}
