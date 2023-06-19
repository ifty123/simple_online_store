package seeder

import (
	"log"
	"time"

	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

func transactionDetailSeeder(db *gorm.DB) {
	now := time.Now()
	var transactionDetail = []model.TransactionDetail{
		{
			TransactionId: 1,
			ProductId:     1,
			Common:        model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
		},
	}

	if err := db.Create(&transactionDetail).Error; err != nil {
		log.Printf("cannot seed data transaction details, with error %v\n", err)
	}
	log.Println("success seed data transaction details")
}
