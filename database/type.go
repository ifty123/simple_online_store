package database

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mysqlConfig struct {
	Host string
	User string
	Pass string
	Port string
	Name string
}

type mongoConfig struct {
	Host string
	User string
	Pass string
	Port string
	Name string
}

func (conf mysqlConfig) Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	var err error

	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func (conf mongoConfig) Connect() *mongo.Database {
	dsn := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.TODO())

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		panic(err)
	}

	dbMongo = client.Database(conf.Name)

	return client.Database(conf.Name)
}
