package seeder

import (
	"github.com/ifty123/simple_online_store/database"
	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{database.GetConnection()}
}

func (s *seed) SeedAll() {
	userSeeder(s.DB)
	categorySeeder(s.DB)
	productSeeder(s.DB)
	cartSeeder(s.DB)
	transactionSeeder(s.DB)
	transactionDetailSeeder(s.DB)
}

func (s *seed) DeleteAll() {
	s.DB.Exec("DELETE FROM users")
	s.DB.Exec("DELETE FROM products")
	s.DB.Exec("DELETE FROM categories")
	s.DB.Exec("DELETE FROM carts")
	s.DB.Exec("DELETE FROM transactions")
	s.DB.Exec("DELETE FROM transaction_details")
}
