package service

import (
	"time"

	model "github.com/polnoy/go-shop/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ShopService is all services
type ShopService interface {
	Gets() ([]model.Shop, error)
	Get(_id string) (model.Shop, error)
	Create(data model.Shop) error
	Update(_id string, data model.Shop) error
	DeleteByID(_id string) error
}

// ShopDb is mongodb
type ShopDb struct {
	ConnectionDB *mgo.Database
}

// ShopCollection is mongo collection name
const ShopCollection = "shops"

// Gets is find all
func (db ShopDb) Gets() ([]model.Shop, error) {
	data := model.Shops{}
	err := db.ConnectionDB.C(ShopCollection).Find(bson.M{}).All(&data)
	return data, err
}

// Get is find once
func (db ShopDb) Get(_id string) (model.Shop, error) {
	data := model.Shop{}
	objectID := bson.ObjectIdHex(_id)
	err := db.ConnectionDB.C(ShopCollection).FindId(objectID).One(&data)
	return data, err
}

// Create is create data
func (db ShopDb) Create(data model.Shop) error {
	data.CreatedTime = time.Now()
	data.UpdatedTime = data.CreatedTime
	return db.ConnectionDB.C(ShopCollection).Insert(data)
}

// Update is update data
func (db ShopDb) Update(_id string, data model.Shop) error {
	objectID := bson.ObjectIdHex(_id)
	newData := bson.M{
		"$set": bson.M{
			"name":         data.Name,
			"detail":       data.Detail,
			"updated_time": time.Now(),
		},
	}
	return db.ConnectionDB.C(ShopCollection).UpdateId(objectID, newData)
}

// DeleteByID is delete by id
func (db ShopDb) DeleteByID(_id string) error {
	objectID := bson.ObjectIdHex(_id)
	return db.ConnectionDB.C(ShopCollection).RemoveId(objectID)
}
