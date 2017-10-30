package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/Gawainli/funnel/utils"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	UID       int           `json:"uid" bson:"uid"`
	UserName  string        `json:"username" bson:"username"`
	Pwd       string        `json:"pwd" bson:"pwd"`
	Accid     string        `json:"accid" bson:"accid"`
	LoginTime time.Time     `json:"logintime" bson:"logintime"`
}

func (a *Account) GenObjectID() bson.ObjectId {
	a.ID = bson.NewObjectId()
	return a.ID
}

func (a *Account) GetObjectID() bson.ObjectId {
	return a.ID
}

func connectMongo() error {
	return ConnectMongoDB("mongodb://192.168.1.111:27017", "test")
}

func Test_MongoDBInsert(t *testing.T) {
	err := connectMongo()
	if err != nil {
		t.Error(err.Error())
	}
	acc := Account{}
	acc.GenObjectID()
	acc.UID = GetNextSeq("testc", "countid", utils.GenRandIntN(100))
	acc.LoginTime = time.Now()
	acc.UserName = "ht5"
	acc.Pwd = "1230000"

	// err = AddData("accounts", js)
	query := func(c *mgo.Collection) error {
		return c.Insert(&acc)
	}
	QueryWithCollection("testaccount", query)
	if err != nil {
		t.Error(err.Error())
	}
}

func Test_MongoDBGetByID(t *testing.T) {
	err := connectMongo()
	if err != nil {
		t.Error(err.Error())
	}

	acc := Account{}
	objid := bson.ObjectIdHex("5007e1e773ea8212f8d0d9d7")
	query := func(c *mgo.Collection) error {
		return c.FindId(objid).One(&acc)
	}

	err = QueryWithCollection("testaccount", query)
	fmt.Println("err:", err.Error())
	fmt.Println("account id", acc.ID.Hex())
	fmt.Println("account time", acc.LoginTime.Local())
}

func Test_MongoDBGet(t *testing.T) {
	err := connectMongo()
	if err != nil {
		t.Error(err.Error())
	}

	acc := Account{}
	key := "username"
	param := "tw_892644980644208640"
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{key: param}).One(&acc)
	}

	err = QueryWithCollection("accounts", query)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("account id", acc.ID.Hex())
	fmt.Println("account time", acc.LoginTime.Local())
}

func Test_MongoDBIncrID(t *testing.T) {
	err := connectMongo()
	if err != nil {
		t.Error(err.Error())
	}
	id := GetNextSeq("testc", "countid", 100)
	fmt.Println(id)
}

func Test_SetSeqBase(t *testing.T) {
	err := connectMongo()
	if err != nil {
		t.Error(err.Error())
	}
	num := SetNextSeqBaseNum("testc", "countid", 1000)
	fmt.Println(num)
}
