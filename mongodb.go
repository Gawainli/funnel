package funnel

import (
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
	defer session.Close()
	c := session.DB(dbConfig.DBName).C(collection)
	return s(c)
}
