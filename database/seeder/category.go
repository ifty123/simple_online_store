package seeder

import (
	"log"
	"time"

	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

func categorySeeder(db *gorm.DB) {
	now := time.Now()
	var category = []model.Category{
		{
			NameCategory: "Sepatu",
			Common:       model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
		},
	}

	if err := db.Create(&category).Error; err != nil {
		log.Printf("cannot seed data category, with error %v\n", err)
	}
	log.Println("success seed data category")
}
