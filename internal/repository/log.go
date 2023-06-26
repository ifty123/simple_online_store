package repository

import (
	"context"
	"log"
	"time"

	"github.com/ifty123/simple_online_store/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepository interface {
	GetAll(ctx context.Context) ([]model.Log, error)
	SaveLog(ctx context.Context, id, payload string, userId uint) error
}

type Log struct {
	Db *mongo.Database
}

func NewLogRepository(db *mongo.Database) *Log {
	return &Log{
		Db: db,
	}
}

func (e *Log) GetAll(ctx context.Context) ([]model.Log, error) {
	var logs []model.Log

	log, err := e.Db.Collection("log").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for log.Next(ctx) {
		var logItem model.Log
		err := log.Decode(&logItem)
		if err != nil {
			return nil, err
		}

		logs = append(logs, model.Log{
			ID:        logItem.ID,
			Payload:   logItem.Payload,
			UserId:    logItem.UserId,
			CreatedAt: logItem.CreatedAt,
		})
	}

	return logs, err
}

func (e *Log) SaveLog(ctx context.Context, id, payload string, userId uint) error {

	saveModel := &model.Log{
		ID:        id,
		Payload:   payload,
		UserId:    userId,
		CreatedAt: time.Now(),
	}

	_, err := e.Db.Collection(model.Log{}.TableName()).InsertOne(ctx, saveModel)
	if err != nil {
		log.Println("err ?", err)
		return err
	}

	return nil
}
