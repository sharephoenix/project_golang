package mgodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestMgo_Connect(t *testing.T) {
	var mgo = &Mgo{
		"mongodb://localhost:27017",
		"user0",
		"table_name0",
		nil,
	}

	mgo.Connect()
}

func TestMgo_Update(t *testing.T) {
	var mgo = &Mgo{
		"mongodb://localhost:27017",
		"user0",
		"table_name0",
		nil,
	}

	mgo.Connect()
	usr := Us{
		"ericn",
		12200,
	}
	mgo.Update(bson.M{"name": "ericluan"}, usr)
}

func TestMgo_FindOne(t *testing.T) {
	var mgo = &Mgo{
		"mongodb://localhost:27017",
		"user0",
		"table_name0",
		nil,
	}

	mgo.Connect()

	var users Us

	err := mgo.FindOne(bson.M{"name": "ericn"}, &users)
	if err != nil {
		fmt.Println("FindOne_Error", err.Error())
	}

	fmt.Println("findOne", users)
}

func TestMgo_FindAll(t *testing.T) {
	var mgo = &Mgo{
		"mongodb://localhost:27017",
		"user",
		"table_name",
		nil,
	}
	mgo.Connect()

	var users []Us

	err := mgo.All(bson.M{"name": "eric"}, &users)
	if err != nil {
		fmt.Println("FindOne_Error", err.Error())
	}

	fmt.Println("findAll", users)
}

func TestMgo_Delete(t *testing.T) {
	var mgo = &Mgo{
		"mongodb://localhost:27017",
		"user",
		"table_name",
		nil,
	}
	mgo.Connect()

	err := mgo.Delete(bson.M{"name": "eric"})
	if err != nil {
		fmt.Println("DELETE_Error", err.Error())
	}

	fmt.Println("DELETE")
}
