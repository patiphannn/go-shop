package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	mongo "github.com/polnoy/go-shop/config"
	model "github.com/polnoy/go-shop/models"
	routes "github.com/polnoy/go-shop/routes"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

func initMain() (*mgo.Database, *gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	os.Setenv("MONGO_DB_NAME", "go-shop-test")

	connectionDB, err := mongo.GetMongoDB()
	router := routes.SetupRouter(connectionDB)

	return connectionDB, router, err
}

var connectionDB, router, errMain = initMain()

func TestInitMain(t *testing.T) {
	assert.NotNil(t, connectionDB)
	assert.NotNil(t, router)
	assert.Nil(t, errMain)
}

// GET /ping API
func TestPingAPI(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["message"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "pong", value)
}

// GET /api/shop API
func TestGetShopsAPI(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/shop", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	response := make(map[string][]model.Shop)
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["shops"]
	len := len(value)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, 0, len)
}

// POST /api/shop API
func TestCreateShopAPI(t *testing.T) {
	var jsonStr = []byte(`{"_id":"5e776f17c8d00a611ed6550f","name":"new test","detail":"new detail"}`)
	req, _ := http.NewRequest("POST", "/api/shop", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response map[string]bool
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["status"]

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, true, value)
}

// PUT /api/shop/:_id API
func TestUpdateShopAPI(t *testing.T) {
	var jsonStr = []byte(`{"name":"update name","detail":"update detail"}`)
	req, _ := http.NewRequest("PUT", "/api/shop/5e776f17c8d00a611ed6550f", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response map[string]bool
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["status"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, true, value)
}

// GET /api/shop/:_id API
func TestGetShopAPI(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/shop/5e776f17c8d00a611ed6550f", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	response := make(map[string]model.Shop)
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["shop"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "update name", value.Name)
	assert.Equal(t, "update detail", value.Detail)
}

// DELETE /api/shop/:_id API
func TestDeletesShopAPI(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/shop/5e776f17c8d00a611ed6550f", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
