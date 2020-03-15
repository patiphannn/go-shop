package service

import (
	model "github.com/polnoy/go-shop/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ShopService interface {
	Gets() ([]model.Shop, error)
	Get(id string) (model.Shop, error)
}

type ShopDb struct {
	ConnectionDB *mgo.Database
}

// Static Collection
const ShopCollection = "shop"

func (db ShopDb) Gets() ([]model.Shop, error) {
	data := model.Shops{}
	err := db.ConnectionDB.C(ShopCollection).Find(bson.M{}).All(&data)
	return data, err
}

func (db ShopDb) Get(id string) (model.Shop, error) {
	data := model.Shop{}
	objectID := bson.ObjectIdHex(id)
	err := db.ConnectionDB.C(ShopCollection).FindId(objectID).One(&data)
	return data, err
}
