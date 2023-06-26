package database

import (
	"github.com/ifty123/simple_online_store/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"

	"gorm.io/gorm"
)

var (
	dbConn  *gorm.DB
	dbMongo *mongo.Database
)

func CreateConnection() {
	conf := mysqlConfig{
		User: util.GetEnv("DB_USER", "root"),
		Pass: util.GetEnv("DB_PASS", "root123"),
		Host: util.GetEnv("DB_HOST", "localhost"),
		Port: util.GetEnv("DB_PORT", "3306"),
		Name: util.GetEnv("DB_NAME", "simple_store"),
	}

	conf.Connect()
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}

//connection mongo (log)
func CreateConnectionLog() {

	var db *mongo.Database

	conf := mongoConfig{
		User: util.GetEnv("MONGO_USER", "root"),
		Pass: util.GetEnv("MONGO_PASS", "root123"),
		Host: util.GetEnv("DB_HOST", "localhost"),
		Port: util.GetEnv("MONGO_PORT", "27017"),
		Name: util.GetEnv("MONGO_NAME", "store_log"),
	}

	db = conf.Connect()

	dbMongo = db
}

func GetConnectionLog() *mongo.Database {

	if dbMongo == nil {
		CreateConnectionLog()
	}

	return dbMongo
}
