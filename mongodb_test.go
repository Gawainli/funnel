package funnel

import (
	"fmt"
	"testing"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	UID      int64         `json:"uid" bson:"uid"`
	UserName string        `json:"username" bson:"uername"`
	Pwd      string        `json:"pwd" bson:"pwd"`
	Accid    string        `json:"accid" bson:"accid"`
}

func (a *Account) GenObjectID() bson.ObjectId {
	a.ID = bson.NewObjectId()
	return a.ID
}

func (a *Account) GetObjectID() bson.ObjectId {
	return a.ID
}

func connectMongo() error {
	return ConnectMongoDB("mongodb://192.168.0.98:27017/test", "test")
}

func Test_MongoDBInsert(t *testing.T) {
	err := connectMongo()
	if err != nil {
		t.Error(err.Error())
	}
	acc := Account{}
	acc.GenObjectID()
	acc.UID = time.Now().UnixNano()
	acc.UserName = "ht223"
	acc.Pwd = "1230000"

	// err = AddData("accounts", js)
	query := func(c *mgo.Collection) error {
		return c.Insert(&acc)
	}
	QueryWithCollection("accounts", query)
	if err != nil {
		t.Error(err.Error())
	}
}

func Test_MongoDBGetByID(t *testing.T) {
	err := connectMongo()
	if err != nil {
		t.Error(err.Error())
	}
	if err != nil {
		t.Error(err.Error())
	}

	acc := Account{}
	objid := bson.ObjectIdHex("5985974c73ea8237186eab52")
	query := func(c *mgo.Collection) error {
		return c.FindId(objid).One(&acc)
	}

	QueryWithCollection("accounts", query)
	fmt.Println("account id", acc.ID.Hex())
}
