package main

import (
	"log"

	mongo "github.com/polnoy/go-shop/config"
	routes "github.com/polnoy/go-shop/routes"
)

func main() {
	connectionDB, err := mongo.GetMongoDB()
	if err != nil {
		log.Panic("Can no connect Database", err.Error())
	}

	r := routes.SetupRouter(connectionDB)
	r.Run(":8080")
}
