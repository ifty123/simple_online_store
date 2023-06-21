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
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewUserRepository(db),
		repository.NewProductRepository(db),
		repository.NewCartRepository(db),
		repository.NewTransactionRepository(db),
		repository.NewTransactionDetailsRepository(db),
	}
}
