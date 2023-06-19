package repository

import (
	"context"

	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

type CartRepository interface {
	UpdateQuantity(ctx context.Context, id uint, quantity int) (*model.Cart, error)
	FindByUserId(ctx context.Context, userId uint) (*model.Cart, error)
	Destroy(ctx context.Context, cart *model.Cart) (*model.Cart, error)
}

type cart struct {
	Db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cart {
	return &cart{
		Db: db,
	}
}

func (e *cart) UpdateQuantity(ctx context.Context, id uint, quantity int) (*model.Cart, error) {
	var prd model.Cart
	q := e.Db.WithContext(ctx).Model(&model.Product{}).Where("id = ?", id)

	err := q.UpdateColumn("quantity", quantity).Error
	return &prd, err
}

func (e *cart) FindByUserId(ctx context.Context, userId uint) (*model.Cart, error) {
	var prd model.Cart
	q := e.Db.WithContext(ctx).Model(&model.Product{}).Where("user_id = ?", userId)

	q = q.Preload("Product")

	err := q.First(&prd).Error
	return &prd, err
}

func (e *cart) Destroy(ctx context.Context, cart *model.Cart) (*model.Cart, error) {
	if err := e.Db.WithContext(ctx).Delete(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
