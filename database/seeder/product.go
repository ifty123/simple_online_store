package seeder

import (
	"log"
	"time"

	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

func productSeeder(db *gorm.DB) {
	now := time.Now()
	var products = []model.Product{
		{
			NameProduct:  "Nike Air",
			PriceProduct: 120000,
			Common:       model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
			CategoryId:   1,
		},
		{
			NameProduct:  "Sport Air",
			PriceProduct: 100000,
			Common:       model.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
			CategoryId:   1,
		},
	}

	if err := db.Create(&products).Error; err != nil {
		log.Printf("cannot seed data product, with error %v\n", err)
	}
	log.Println("success seed data product")
}
