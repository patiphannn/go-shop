package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/polnoy/go-shop/controllers"
	service "github.com/polnoy/go-shop/services"
	"gopkg.in/mgo.v2"
)

func SetupRouter(connectionDB *mgo.Database) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	shopService := service.ShopDb{
		ConnectionDB: connectionDB,
	}

	shopController := controller.ShopAPI{
		ShopService: &shopService,
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.GET("/shops", shopController.Gets)
	api.GET("/shop/:id", shopController.Get)

	return r
}
