package mgodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

func TestMgo_Connect(t *testing.T) {
	var mgo = &Mgo{
		"mongodb://127.0.0.1:31017",
		"user0",
		"table_name0",
		nil,
	}

	mgo.Connect()
}

func TestMgo_InsertOne(t *testing.T) {
	var collection mongo.Collection
	var mgo = &Mgo{
		"mongodb://127.0.0.1:31017",
		"user0",
		"table_name0",
		&collection,
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mgo.Uri))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// 插入相关数据
	collection0 := client.Database(mgo.Database).Collection(mgo.Collection)
	fmt.Println("beginning_insert", ctx)
	res, err := collection0.InsertOne(context.Background(), bson.M{"name": "alexluan", "age": 122})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	id := res.InsertedID
	fmt.Println("id:", id, res)

	// 找到相关数据
	//var resMap map[string]interface{}
	//cur, err := collection0.Find(context.Background(), bson.M{"name": "alexluan"})
	//cur.Next(context.Background())
	//fmt.Println(err, resMap, cur.Decode(cur.Current))
}

func TestMgo_Update(t *testing.T) {
	var collection mongo.Collection
	var mgo = &Mgo{
		"mongodb://localhost:31017",
		"user0",
		"table_name0",
		&collection,
	}
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	//defer cancel() //养成良好的习惯，在调用WithTimeout之后defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mgo.Uri))
	if err != nil {
		log.Print(err)
	}
	collection0 := client.Database(mgo.Database).Collection(mgo.Collection)
	fmt.Println("lll:", collection0)
	mgo.MgoCollection = collection0
	usr := Us{
		"ericn",
		12200,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	result, err := mgo.MgoCollection.UpdateOne(ctx, bson.M{"name": "alexluan"}, bson.M{"$set": usr})
	if err != nil {
		fmt.Println("修改失败", err)
	} else {
		fmt.Println("修改成功", result)
	}
	//mgo.Update(bson.M{"name": "alexluan"}, usr)
}

func TestMgo_FindOne(t *testing.T) {
	var mgo = &Mgo{
		"mongodb://127.0.0.1:31017",
		"user0",
		"table_name0",
		nil,
	}

	mgo.Connect()

	var users Us

	err := mgo.FindOne(bson.M{"name": "alexluan"}, &users)
	if err != nil {
		fmt.Println("FindOne_Error", err.Error())
	}

	fmt.Println("findOne", users.Name)
}

func TestMgo_FindAll(t *testing.T) {
	var mgo = &Mgo{
		"mongodb://127.0.0.1:31017",
		"user0",
		"table_name0",
		nil,
	}
	mgo.Connect()

	var users map[string]interface{}

	err := mgo.All(bson.M{}, users)
	if err != nil {
		fmt.Println("FindOne_Error", err.Error())
		return
	}

	fmt.Println("findAll", users)
}

func TestMgo_Delete(t *testing.T) {
	var mgo = &Mgo{
		"mongodb://localhost:31017",
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
