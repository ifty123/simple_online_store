package model

import "time"

type Log struct {
	ID        string    `bson:"_id"`
	Payload   string    `bson:"payload"`
	UserId    uint      `bson:"user_id"`
	CreatedAt time.Time `bson:"created_at"`
}

func (Log) TableName() string {
	return "log"
}
