package kvdb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"zest/engine/common"
	"zest/engine/conf"
	"zest/engine/zslog"
)

type mongoConf struct {
	ip       string
	port     int
	user     string
	password string
	poolSize int
	timeout  int
	idleTime int

	uri string
}

const (
	MONGO_IP        = "127.0.0.1"
	MONGO_PORT      = 27017
	MONGO_USER      = "admin"
	MONGO_PW        = "password"
	MONGO_POOL_SIZE = 1
	MONGO_TIMEOUT   = 5000
	MONGO_IDLE_TIME = 5000
)

var (
	MongoClient   *mongo.Client
	mongoConfInfo *mongoConf
)

func init() {
	initConfig()
}

func initConfig() {
	mongoConfInfo = new(mongoConf)
	ip := common.ThreeUnary(conf.IsSet("mongo.ip"), conf.GetString("mongo.ip"), MONGO_IP)
	port := common.ThreeUnary(conf.IsSet("mongo.port"), conf.GetInt("mongo.port"), MONGO_PORT)
	user := common.ThreeUnary(conf.IsSet("mongo.user"), conf.GetString("mongo.user"), MONGO_USER)
	password := common.ThreeUnary(conf.IsSet("mongo.password"), conf.GetString("mongo.password"), MONGO_PW)
	poolSize := common.ThreeUnary(conf.IsSet("mongo.poolSize"), conf.GetInt("mongo.poolSize"), MONGO_POOL_SIZE)
	timeout := common.ThreeUnary(conf.IsSet("mongo.timeout"), conf.GetInt("mongo.timeout"), MONGO_TIMEOUT)
	idleTime := common.ThreeUnary(conf.IsSet("mongo.idleTime"), conf.GetInt("mongo.idleTime"), MONGO_IDLE_TIME)
	mongoConfInfo.ip = common.String(ip)
	mongoConfInfo.port = common.Int(port)
	mongoConfInfo.user = common.String(user)
	mongoConfInfo.password = common.String(password)
	mongoConfInfo.poolSize = common.Int(poolSize)
	mongoConfInfo.timeout = common.Int(timeout)
	mongoConfInfo.idleTime = common.Int(idleTime)
	mongoConfInfo.generateMongoURI()
}

func (mc *mongoConf) generateMongoURI() {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v", mc.user, mc.password, mc.ip, mc.port)
	mc.uri = uri
}

func Connect() {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoConfInfo.uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		zslog.LogError("Connect MongoDB Error %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		zslog.LogError("Ping MongoDB Error %v", err)
	}

	zslog.LogDebug("Connected to MongoDB!")
	MongoClient = client
	return
}

func Close() {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		zslog.LogError("Disconnect MongoDB Error %v", err)
	}
	zslog.LogInfo("Disconnect MongoDB Success!")
}

func InsertOne(dbname string, table string, doc interface{}) {
	coll := MongoClient.Database(dbname).Collection(table)
	coll.InsertOne(context.TODO(), doc)
}

// func FindOne(){

// }

// func FindOne(dbname string,table string,cond)
