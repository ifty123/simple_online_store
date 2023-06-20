package repository

import (
	"context"

	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTransacion(ctx context.Context, payload *dto.TransactionReq) (*model.Transaction, error)
}

type transaction struct {
	Db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transaction {
	return &transaction{
		Db: db,
	}
}

func (e *transaction) SaveTransacion(ctx context.Context, payload *dto.TransactionReq) (*model.Transaction, error) {
	//save ke transaction
	newTransaction := model.Transaction{
		UserId:            payload.UserId,
		TotalTransaction:  payload.Total,
		StatusTransaction: dto.BELUM_DIBAYAR,
	}

	if err := e.Db.WithContext(ctx).Save(&newTransaction).Error; err != nil {
		return &newTransaction, err
	}

	//save ke transaction details
	for _, i := range payload.Cart {
		newDetails := model.TransactionDetail{
			TransactionId: newTransaction.ID,
			ProductId:     i.ProductId,
			Quantity:      i.Quantity,
			PriceTotal:    i.Price,
		}

		if err := e.Db.WithContext(ctx).Save(&newDetails).Error; err != nil {
			return &newTransaction, err
		}
	}
	return &newTransaction, nil

}
