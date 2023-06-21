package repository

import (
	"context"

	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, payload *dto.TransactionReq) (*model.Transaction, error)
	FindByUserId(ctx context.Context, userId uint) ([]model.Transaction, error)
	FindById(ctx context.Context, userId uint) (model.Transaction, error)
	UpdateTransaction(ctx context.Context, id uint, payload *model.Transaction) (*model.Transaction, error)
}

type transaction struct {
	Db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transaction {
	return &transaction{
		Db: db,
	}
}

func (e *transaction) SaveTransaction(ctx context.Context, payload *dto.TransactionReq) (*model.Transaction, error) {
	//save ke transaction
	newTransaction := model.Transaction{
		UserId:            payload.UserId,
		TotalTransaction:  payload.Total,
		StatusTransaction: dto.BELUM_DIBAYAR,
	}

	if err := e.Db.WithContext(ctx).Save(&newTransaction).Error; err != nil {
		return &newTransaction, err
	}

	return &newTransaction, nil

}

func (e *transaction) FindByUserId(ctx context.Context, userId uint) ([]model.Transaction, error) {
	var prd []model.Transaction
	q := e.Db.WithContext(ctx).Model(&model.Transaction{}).Where("user_id = ?", userId)

	err := q.Find(&prd).Error
	return prd, err
}

func (e *transaction) UpdateTransaction(ctx context.Context, id uint, payload *model.Transaction) (*model.Transaction, error) {

	result := e.Db.WithContext(ctx).Model(&model.Transaction{}).Where("id = ?", id).Updates(&payload)
	if result.Error != nil {
		return nil, result.Error
	}
	return payload, nil
}

func (e *transaction) FindById(ctx context.Context, userId uint) (model.Transaction, error) {
	var prd model.Transaction
	q := e.Db.WithContext(ctx).Model(&model.Transaction{}).Where("id = ?", userId)

	err := q.First(&prd).Error
	return prd, err
}
