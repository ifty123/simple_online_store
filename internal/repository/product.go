package repository

import (
	"context"

	"github.com/ifty123/simple_online_store/internal/model"
	pkgdto "github.com/ifty123/simple_online_store/pkg/dto"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByCategory(ctx context.Context, category uint, pagination *pkgdto.Pagination) ([]model.Product, *pkgdto.PaginationInfo, error)
	FindById(ctx context.Context, id uint, categoryPreload bool) (*model.Product, error)
	FindByIds(ctx context.Context, ids []uint) ([]model.Product, error)
}

type product struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *product {
	return &product{
		Db: db,
	}
}

func (e *product) FindByCategory(ctx context.Context, category uint, pagination *pkgdto.Pagination) ([]model.Product, *pkgdto.PaginationInfo, error) {
	var user []model.Product
	var count int64

	query := e.Db.WithContext(ctx).Model(&model.Product{})

	if category != 0 {
		query = query.Where("category_id = ?", category)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&user).Error

	return user, pkgdto.CheckInfoPagination(pagination, count), err
}

func (e *product) FindById(ctx context.Context, id uint, categoryPreload bool) (*model.Product, error) {
	var prd model.Product
	q := e.Db.WithContext(ctx).Model(&model.Product{}).Where("id = ?", id)

	if categoryPreload {
		q = q.Preload("Category")
	}

	err := q.First(&prd).Error
	return &prd, err
}

func (e *product) FindByIds(ctx context.Context, ids []uint) ([]model.Product, error) {
	var prd []model.Product
	q := e.Db.WithContext(ctx).Model(&model.Product{}).Where("id IN (?)", ids)

	err := q.Find(&prd).Error
	return prd, err
}
