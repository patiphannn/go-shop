package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/polnoy/go-shop/services"
)

type ShopAPI struct {
	ShopService service.ShopService
}

func (api ShopAPI) Gets(c *gin.Context) {
	data, err := api.ShopService.Gets()
	if err != nil {
		log.Println("error shop Gets: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"shops": &data,
	})
}

func (api ShopAPI) Get(c *gin.Context) {
	id := c.Param("id")
	data, err := api.ShopService.Get(id)
	if err != nil {
		log.Println("error shop Get", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"shop": &data,
	})
}
