package seeder

import (
	"log"
	"time"

	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

func userSeeder(db *gorm.DB) {
	now := time.Now()
	var users = []model.User{
		{
			Username: "Alifipa5",
			Email:    "iftAlif@gmail.com",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			Common:   model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
		},
		{
			Username: "crystal_jung",
			Email:    "crystal@gmail.com",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			Common:   model.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
		},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Printf("cannot seed data employees, with error %v\n", err)
	}
	log.Println("success seed data employees")
}
