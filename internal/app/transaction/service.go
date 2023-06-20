package transaction

import (
	"context"
	"log"

	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/factory"
	"github.com/ifty123/simple_online_store/internal/repository"

	"github.com/ifty123/simple_online_store/pkg/util/response"
)

type Service struct {
	TrasactionRepository repository.TransactionRepository
	ProductRepository    repository.ProductRepository
	CartRepository       repository.CartRepository
}

type CartService interface {
	SaveTransaction(ctx context.Context, payload *dto.TransactionReq) (*dto.TransactionResponse, error)
}

func Newservice(f *factory.Factory) Service {
	return Service{
		TrasactionRepository: f.TransactionRepository,
		ProductRepository:    f.ProductRepository,
		CartRepository:       f.CartRepository,
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

	var productId []uint
	for _, i := range cart {
		productId = append(productId, i.ProductId)
		productRes[i.ProductId] = &dto.ProductDetailResponse{
			ID:         i.ProductId,
			PriceTotal: i.PriceTotal,
			Quantity:   i.Quantity,
		}

		payload.Total += i.PriceTotal
	}

	//find product by id
	product, err := s.ProductRepository.FindByIds(ctx, productId)
	if err != nil {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	for _, i := range product {
		log.Println("nama produk :", i.NameProduct, " - id :", i.ID)
		if _, ok := productRes[i.ID]; ok {
			productRes[i.ID].NameProduct = i.NameProduct
			productRes[i.ID].Price = i.PriceProduct
		}

		productArrRes = append(productArrRes, productRes[i.ID])
	}

	//save transation
	transaction, err := s.TrasactionRepository.SaveTransacion(ctx, payload)
	if err != nil {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	res = &dto.TransactionResponse{
		ID:                transaction.ID,
		Product:           productArrRes,
		Price:             transaction.TotalTransaction,
		StatusTransaction: transaction.StatusTransaction,
	}

	return res, nil
}

/*
func (s *Service) FindCartByUserId(ctx context.Context, userId uint) ([]*dto.CartResponse, error) {

	cart, err := s.CartRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	productRes := []*dto.CartResponse{}
	if len(cart) != 0 {
		for _, i := range cart {

			product := dto.ProductResponse{
				ID:          i.Product.ID,
				NameProduct: i.Product.NameProduct,
				Price:       i.Product.PriceProduct,
			}

			productRes = append(productRes, &dto.CartResponse{
				ID:       i.ID,
				Product:  product,
				Quantity: int64(i.Quantity),
				Price:    int64(i.PriceTotal),
			})
		}
	}

	return productRes, nil
}

*/
