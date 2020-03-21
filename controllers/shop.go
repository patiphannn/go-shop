package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/polnoy/go-shop/models"
	service "github.com/polnoy/go-shop/services"
)

// ShopAPI is ShopService
type ShopAPI struct {
	ShopService service.ShopService
}

// Gets is find all
func (api ShopAPI) Gets(c *gin.Context) {
	data, err := api.ShopService.Gets()
	if err != nil {
		log.Println("error shop Gets: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shops": &data,
	})
}

// Get is find once
func (api ShopAPI) Get(c *gin.Context) {
	_id := c.Param("_id")
	data, err := api.ShopService.Get(_id)
	if err != nil {
		log.Println("error shop Get", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shop": &data,
	})
}

// Create is create data
func (api ShopAPI) Create(c *gin.Context) {
	data := model.Shop{}
	err := c.ShouldBindJSON(&data)

	if err != nil {
		log.Println("error shop Create", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = api.ShopService.Create(data)

	if err != nil {
		log.Println("error shop Create", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
	})
}

// Update is update data
func (api ShopAPI) Update(c *gin.Context) {
	_id := c.Param("_id")
	data := model.Shop{}
	err := c.ShouldBindJSON(&data)

	if err != nil {
		log.Println("error shop Update", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = api.ShopService.Update(_id, data)

	if err != nil {
		log.Println("error shop Update", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

// DeleteByID is delete by id
func (api ShopAPI) DeleteByID(c *gin.Context) {
	_id := c.Param("_id")
	err := api.ShopService.DeleteByID(_id)

	if err != nil {
		log.Println("error shop Delete", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{
		"status": true,
	})
}
