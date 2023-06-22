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
			NameProduct:  "Sneakers Pria",
			PriceProduct: 100000,
			Common:       model.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
			CategoryId:   1,
		},
		{
			NameProduct:  "Kemeja Denim",
			PriceProduct: 80000,
			Common:       model.Common{ID: 3, CreatedAt: now, UpdatedAt: now},
			CategoryId:   2,
		},
		{
			NameProduct:  "Kemeja Lengan Panjang Ufaira",
			PriceProduct: 80000,
			Common:       model.Common{ID: 4, CreatedAt: now, UpdatedAt: now},
			CategoryId:   2,
		},
		{
			NameProduct:  "Topi Hitam",
			PriceProduct: 50000,
			Common:       model.Common{ID: 5, CreatedAt: now, UpdatedAt: now},
			CategoryId:   3,
		},
	}

	if err := db.Create(&products).Error; err != nil {
		log.Printf("cannot seed data product, with error %v\n", err)
	}
	log.Println("success seed data product")
}
