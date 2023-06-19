package factory

import (
	"github.com/ifty123/simple_online_store/database"
	"github.com/ifty123/simple_online_store/internal/repository"
)

type Factory struct {
	UserRepository repository.UserRepository
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewUserRepository(db),
	}
}
