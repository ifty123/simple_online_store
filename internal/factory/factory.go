package factory

import (
	"github.com/ifty123/simple_online_store/database"
	"github.com/ifty123/simple_online_store/internal/repository"
)

type Factory struct {
	UserRepository               repository.UserRepository
	ProductRepository            repository.ProductRepository
	CartRepository               repository.CartRepository
	TransactionRepository        repository.TransactionRepository
	TransactionDetailsRepository repository.TransactionDetailsRepository
	CategoryRepository           repository.CategoryRepository
	LogRepository                repository.LogRepository
}

func NewFactory() *Factory {
	db := database.GetConnection()
	dbMongo := database.GetConnectionLog()
	return &Factory{
		repository.NewUserRepository(db),
		repository.NewProductRepository(db),
		repository.NewCartRepository(db),
		repository.NewTransactionRepository(db),
		repository.NewTransactionDetailsRepository(db),
		repository.NewcategoryRepository(db),
		repository.NewLogRepository(dbMongo),
	}
}
