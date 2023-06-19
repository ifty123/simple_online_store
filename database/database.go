package database

import (
	"sync"

	"github.com/ifty123/simple_online_store/pkg/util"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {
	conf := dbConfig{
		User: util.GetEnv("DB_USER", "root"),
		Pass: util.GetEnv("DB_PASS", "root123"),
		Host: util.GetEnv("DB_HOST", "localhost"),
		Port: util.GetEnv("DB_PORT", "3306"),
		Name: util.GetEnv("DB_NAME", "simple_store"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
