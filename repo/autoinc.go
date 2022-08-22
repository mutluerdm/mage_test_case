package repo

import (
	"context"
	"mage_test_case/mlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	AutoInc struct {
		collection   *mongo.Collection
		idFieldName  string
		seqFieldName string
	}
)

// Connect to database
func NewUtoInc(c *mongo.Collection) *AutoInc {
	return &AutoInc{collection: c, idFieldName: "id", seqFieldName: "seq"}
}

// Next sequence of AutoIncrement
func (ai *AutoInc) Next(name string) uint64 {
	retunDoc := options.After
	filter := bson.M{ai.idFieldName: name}
	upsert := true
	opts := &options.FindOneAndUpdateOptions{ReturnDocument: &retunDoc, Upsert: &upsert}
	res := ai.collection.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": bson.M{ai.idFieldName: name}, "$inc": bson.M{ai.seqFieldName: 1}}, opts)
	if res.Err() != nil {
		mlog.Println("Autoincrement error(1):", res.Err())
	}
	jsonData := &bson.D{}
	_ = res.Decode(jsonData)
	sec := jsonData.Map()["seq"].(int32)
	return uint64(sec)
}

// Cancel is decrement counter value
func (ai *AutoInc) Cancel(name string) {
	_, err := ai.collection.UpdateOne(context.Background(), bson.M{"id": name}, bson.M{"$inc": bson.M{"seq": -1}})
	if err != nil {
		mlog.Println("Autoincrement error(2):", err)
	}
}
