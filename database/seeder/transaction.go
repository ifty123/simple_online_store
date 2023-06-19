package seeder

import (
	"log"
	"time"

	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

func transactionSeeder(db *gorm.DB) {
	now := time.Now()
	var transaction = []model.Transaction{
		{
			UserId:            1,
			TotalTransaction:  120000,
			StatusTransaction: "belum dibayar",
			Common:            model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
		},
	}

	if err := db.Create(&transaction).Error; err != nil {
		log.Printf("cannot seed data transaction, with error %v\n", err)
	}
	log.Println("success seed data transaction")
}
