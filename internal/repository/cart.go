package repository

import (
	"context"

	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

type CartRepository interface {
	UpdateCart(ctx context.Context, id uint, payload *model.Cart) (*model.Cart, error)
	FindByUserId(ctx context.Context, userId uint) ([]model.Cart, error)
	Destroy(ctx context.Context, cart *model.Cart) (*model.Cart, error)
	SaveCart(ctx context.Context, cart *dto.Cart) (*model.Cart, error)
	FindByProductId(ctx context.Context, productId uint, userId uint) (*model.Cart, error)
	FindById(ctx context.Context, id uint) (*model.Cart, error)
	FindByIds(ctx context.Context, ids []uint) ([]model.Cart, error)
}

type cart struct {
	Db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cart {
	return &cart{
		Db: db,
	}
}

func (e *cart) UpdateCart(ctx context.Context, id uint, payload *model.Cart) (*model.Cart, error) {

	result := e.Db.WithContext(ctx).Model(&model.Cart{}).Where("id = ?", id).Updates(&payload)
	if result.Error != nil {
		return nil, result.Error
	}
	return payload, nil
}

func (e *cart) FindByUserId(ctx context.Context, userId uint) ([]model.Cart, error) {
	var prd []model.Cart
	q := e.Db.WithContext(ctx).Preload("Product").Model(&model.Cart{}).Where("user_id = ?", userId)

	err := q.Find(&prd).Error
	return prd, err
}

func (e *cart) FindByProductId(ctx context.Context, productId uint, userId uint) (*model.Cart, error) {
	var prd model.Cart
	q := e.Db.WithContext(ctx).Model(&model.Cart{}).Where("product_id = ?", productId).Where("user_id = ?", userId)

	err := q.First(&prd).Error
	return &prd, err
}

func (e *cart) FindById(ctx context.Context, id uint) (*model.Cart, error) {
	var prd model.Cart
	q := e.Db.WithContext(ctx).Model(&model.Cart{}).Where("id = ?", id)

	err := q.First(&prd).Error
	return &prd, err
}

func (e *cart) Destroy(ctx context.Context, cart *model.Cart) (*model.Cart, error) {
	if err := e.Db.WithContext(ctx).Delete(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (e *cart) SaveCart(ctx context.Context, cart *dto.Cart) (*model.Cart, error) {
	newCart := model.Cart{
		UserId:     cart.UserId,
		ProductId:  cart.ProductId,
		Quantity:   cart.Quantity,
		PriceTotal: cart.Price,
	}

	if err := e.Db.WithContext(ctx).Save(&newCart).Error; err != nil {
		return &newCart, err
	}

	return &newCart, nil
}

func (e *cart) FindByIds(ctx context.Context, ids []uint) ([]model.Cart, error) {
	var prd []model.Cart
	q := e.Db.WithContext(ctx).Model(&model.Cart{}).Where("id IN (?)", ids)

	err := q.Find(&prd).Error
	return prd, err
}
