package db

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongodbConfig DB配置
type MongodbConfig struct {
	URL    string
	DBName string
}

//MongoDataInterface 存储的MongoData
type MongoDataInterface interface {
	GenObjectID() bson.ObjectId
	GetObjectID() bson.ObjectId
}

var mgoSession *mgo.Session
var dbConfig *MongodbConfig

//ConnectMongoDB 连接MongoDB
func ConnectMongoDB(url string, dbName string) error {
	dbConfig = &MongodbConfig{}
	dbConfig.URL = url
	dbConfig.DBName = dbName
	var err error
	mgoSession, err = mgo.Dial(url)

	if err != nil {
		return err
	}
	return nil
}

//ConnectMongoDB 连接MongoDB
func ConnectMongoDBWithPasswd(url string, dbName string, user string, passwd string, timeout time.Duration) error {
	if strings.HasPrefix(url, "mongodb://") {
		url = url[10:]
	}
	url = user + ":" + passwd + "@" + url + "/" + dbName
	fmt.Println("url: " + url)
	var err error
	mgoSession, err = mgo.DialWithTimeout(url, timeout)
	//mgoSession.Ping()
	dbConfig = &MongodbConfig{}
	dbConfig.URL = url
	dbConfig.DBName = dbName

	if err != nil {
		return err
	}
	return nil
}

//GetMongoSession 获取session， 必须手动Close
func getMongoSession() *mgo.Session {
	if mgoSession == nil {
		return nil
	}
	//最大连接池默认为4096
	return mgoSession.Clone()
}

//QueryWithCollection 获取collection对象
func QueryWithCollection(collection string, s func(*mgo.Collection) error) error {
	session := getMongoSession()
	if session != nil {
		defer session.Close()
		c := session.DB(dbConfig.DBName).C(collection)
		return s(c)
	}

	return errors.New("session is nil")
}

func SetNextSeqBaseNum(collection string, key string, base int) int {
	doc := struct{ Seq int }{}
	query := func(c *mgo.Collection) error {
		change := mgo.Change{
			Update: bson.M{
				"$setOnInsert": bson.M{"seq": base},
			},
			Upsert:    true,
			ReturnNew: true,
		}
		_, err := c.Find(bson.M{"_id": key}).Apply(change, &doc)
		if err != nil {
			panic(fmt.Errorf("set counter failed:", err.Error()))
		}
		return err
	}
	err := QueryWithCollection(collection, query)
	if err != nil {
		return -1
	}
	return doc.Seq
}

//GetNextSeq 获取自增id
func GetNextSeq(collection string, key string, num int) int {
	doc := struct{ Seq int }{}
	cid := key
	query := func(c *mgo.Collection) error {
		change := mgo.Change{
			Update:    bson.M{"$inc": bson.M{"seq": num}},
			Upsert:    true,
			ReturnNew: true,
		}
		_, err := c.Find(bson.M{"_id": cid}).Apply(change, &doc)
		if err != nil {
			panic(fmt.Errorf("get counter failed:", err.Error()))
		}
		return err
	}

	err := QueryWithCollection(collection, query)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return doc.Seq
}
