package mgodb

import (
	"context"
	"example.com/m/common/baseresponse"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// 仅仅被用于测试
type Us struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type Mgo struct {
	Uri           string //数据库网络地址
	Database      string //要连接的数据库
	Collection    string //要连接的集合
	MgoCollection *mongo.Collection
}

func (m *Mgo) Connect() *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel() //养成良好的习惯，在调用WithTimeout之后defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Uri))
	if err != nil {
		log.Print(err)
	}
	collection := client.Database(m.Database).Collection(m.Collection)
	fmt.Println("lll:", collection)
	m.MgoCollection = collection
	return collection
}

func (m *Mgo) All(filter interface{}, results interface{}) error {
	if filter == nil {
		return &baseresponse.LysError{"查询条件，eg：bson.M{\"name\": \"alex\"}"}
	}

	if results == nil {
		return &baseresponse.LysError{"缺少参数 result"}
	}

	collection := m.MgoCollection

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	cursor.All(ctx, results)
	return nil
}

func (m *Mgo) FindOne(filter interface{}, result interface{}) error {
	collection := m.MgoCollection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	signalResult := collection.FindOne(ctx, filter)
	if signalResult.Err() != nil {
		return &baseresponse.LysError{signalResult.Err().Error()}
	}
	err := signalResult.Decode(result)
	return err
}

func (m *Mgo) InsertOne(data interface{}) error {
	if data == nil {
		return &baseresponse.LysError{"插入数据不正确"}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := m.MgoCollection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

//根据id进行修改
func (m *Mgo) Update(filter, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := m.MgoCollection.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		fmt.Println("修改失败", result.ModifiedCount)
	} else {
		fmt.Println("修改成功", result.ModifiedCount)
	}
	return err
}

func (m *Mgo) Delete(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := m.MgoCollection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Println("删除记录", result.DeletedCount, "条")
	return nil
}
