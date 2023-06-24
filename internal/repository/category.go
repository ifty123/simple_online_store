package repository

import (
	"context"

	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]model.Category, error)
}

type category struct {
	Db *gorm.DB
}

func NewcategoryRepository(db *gorm.DB) *category {
	return &category{
		Db: db,
	}
}

func (e *category) GetAll(ctx context.Context) ([]model.Category, error) {
	var user []model.Category

	query := e.Db.WithContext(ctx).Model(&model.Category{})

	err := query.Find(&user).Error

	return user, err
}
