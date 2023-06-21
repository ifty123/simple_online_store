package transaction

import (
	"context"
	"errors"

	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/factory"
	"github.com/ifty123/simple_online_store/internal/repository"

	"github.com/ifty123/simple_online_store/pkg/util/response"
)

type Service struct {
	TrasactionRepository         repository.TransactionRepository
	ProductRepository            repository.ProductRepository
	CartRepository               repository.CartRepository
	TransactionDetailsRepository repository.TransactionDetailsRepository
}

type CartService interface {
	SaveTransaction(ctx context.Context, payload *dto.TransactionReq) (*dto.TransactionResponse, error)
	FindTransactionByUserId(ctx context.Context, userId uint) ([]*dto.TransactionResponse, error)
	UpdateTransactionById(ctx context.Context, userId, transactionId uint) (*dto.TransactionResponse, error)
}

func Newservice(f *factory.Factory) Service {
	return Service{
		TrasactionRepository:         f.TransactionRepository,
		ProductRepository:            f.ProductRepository,
		CartRepository:               f.CartRepository,
		TransactionDetailsRepository: f.TransactionDetailsRepository,
	}
}

func (s *Service) SaveTransaction(ctx context.Context, payload *dto.TransactionReq) (*dto.TransactionResponse, error) {

	var res *dto.TransactionResponse
	productArrRes := []*dto.ProductDetailResponse{}
	productRes := make(map[uint]*dto.ProductDetailResponse)

	//find product by id
	cart, err := s.CartRepository.FindByIds(ctx, payload.CartId)
	if err != nil {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if len(cart) == 0 {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, errors.New("Cart not found"))
	}

	var productId []uint
	for _, i := range cart {

		if _, ok := productRes[i.ProductId]; !ok && i.ProductId != 0 {
			productId = append(productId, i.ProductId)
			productRes[i.ProductId] = &dto.ProductDetailResponse{
				ID:         i.ProductId,
				PriceTotal: i.PriceTotal,
				Quantity:   i.Quantity,
			}

			payload.Total += i.PriceTotal
		}

	}

	//find product by id
	product, err := s.ProductRepository.FindByIds(ctx, productId)
	if err != nil {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	for _, i := range product {
		if _, ok := productRes[i.ID]; ok {
			productRes[i.ID].NameProduct = i.NameProduct
			productRes[i.ID].Price = i.PriceProduct
		}

		productArrRes = append(productArrRes, productRes[i.ID])
	}

	//save transation
	transaction, err := s.TrasactionRepository.SaveTransaction(ctx, payload)
	if err != nil {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	for _, i := range cart {
		_, err := s.TransactionDetailsRepository.SaveTransactionDetails(ctx, &i, transaction.ID)
		if err != nil {
			return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
	}

	res = &dto.TransactionResponse{
		ID:                transaction.ID,
		Product:           productArrRes,
		Price:             transaction.TotalTransaction,
		StatusTransaction: transaction.StatusTransaction,
	}

	return res, nil
}

func (s *Service) FindTransactionByUserId(ctx context.Context, userId uint) ([]*dto.TransactionResponse, error) {

	transactions, err := s.TrasactionRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	productRes := []*dto.TransactionResponse{}
	if len(transactions) != 0 {

		for _, i := range transactions {

			var productDetail []*dto.ProductDetailResponse

			details, err := s.TransactionDetailsRepository.FindByTransactionId(ctx, i.ID)
			if err != nil {
				return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
			}

			for _, d := range details {
				if d.TransactionId == i.ID {
					productDetail = append(productDetail, &dto.ProductDetailResponse{
						ID:          d.Product.ID,
						NameProduct: d.Product.NameProduct,
						Quantity:    d.Quantity,
					})
				}
			}

			//masukkan ke productRes
			productRes = append(productRes, &dto.TransactionResponse{
				ID:                i.ID,
				Product:           productDetail,
				Price:             i.TotalTransaction,
				StatusTransaction: i.StatusTransaction,
			})
		}
	}

	return productRes, nil
}

func (s *Service) UpdateTransactionById(ctx context.Context, userId, transactionId uint) (*dto.TransactionResponse, error) {

	transactions, err := s.TrasactionRepository.FindById(ctx, transactionId)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	//update status ke bayar
	transactions.StatusTransaction = dto.SUDAH_DIBAYAR

	_, err = s.TrasactionRepository.UpdateTransaction(ctx, transactionId, &transactions)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	res := &dto.TransactionResponse{
		ID:                transactions.ID,
		Price:             transactions.TotalTransaction,
		StatusTransaction: transactions.StatusTransaction,
	}
	return res, nil
}
