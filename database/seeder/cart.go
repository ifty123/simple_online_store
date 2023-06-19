package seeder

import (
	"log"
	"time"

	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

func cartSeeder(db *gorm.DB) {
	now := time.Now()
	var carts = []model.Cart{
		{
			UserId:    1,
			ProductId: 1,
			Quantity:  1,
			Common:    model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
		},
	}

	if err := db.Create(&carts).Error; err != nil {
		log.Printf("cannot seed data carts, with error %v\n", err)
	}
	log.Println("success seed data carts")
}
