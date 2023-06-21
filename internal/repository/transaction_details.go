package repository

import (
	"context"

	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

type TransactionDetailsRepository interface {
	SaveTransactionDetails(ctx context.Context, payload *model.Cart, transactionId uint) (*model.TransactionDetail, error)
	FindByTransactionId(ctx context.Context, transactionId uint) ([]model.TransactionDetail, error)
}

type transactionDetails struct {
	Db *gorm.DB
}

func NewTransactionDetailsRepository(db *gorm.DB) *transactionDetails {
	return &transactionDetails{
		Db: db,
	}
}

func (e *transactionDetails) SaveTransactionDetails(ctx context.Context, payload *model.Cart, transactionId uint) (*model.TransactionDetail, error) {

	//save ke transaction details
	newDetails := model.TransactionDetail{
		TransactionId: transactionId,
		ProductId:     payload.ProductId,
		Quantity:      payload.Quantity,
		PriceTotal:    payload.PriceTotal,
	}

	if err := e.Db.WithContext(ctx).Save(&newDetails).Error; err != nil {
		return &newDetails, err
	}
	return &newDetails, nil

}

func (e *transactionDetails) FindByTransactionId(ctx context.Context, transactionId uint) ([]model.TransactionDetail, error) {
	var prd []model.TransactionDetail
	q := e.Db.WithContext(ctx).Preload("Product").Model(&model.TransactionDetail{}).Where("transaction_id = ?", transactionId)

	err := q.Find(&prd).Error
	return prd, err
}
