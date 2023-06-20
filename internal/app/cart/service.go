package cart

import (
	"context"
	"log"

	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/factory"
	"github.com/ifty123/simple_online_store/internal/repository"

	"github.com/ifty123/simple_online_store/pkg/util/response"
)

type Service struct {
	CartRepository    repository.CartRepository
	ProductRepository repository.ProductRepository
}

type CartService interface {
	FindCartByUserId(ctx context.Context, userId uint) ([]*dto.CartResponse, error)
	SaveCart(ctx context.Context, payload *dto.Cart) (*dto.CartResponse, error)
}

func Newservice(f *factory.Factory) Service {
	return Service{
		CartRepository:    f.CartRepository,
		ProductRepository: f.ProductRepository,
	}
}

func (s *Service) SaveCart(ctx context.Context, payload *dto.Cart) (*dto.CartResponse, error) {

	var res *dto.CartResponse

	cart, err := s.CartRepository.FindByProductId(ctx, payload.ProductId)
	if err != nil {
		log.Println("error no found : ", err)
	}

	//find product
	product, err := s.ProductRepository.FindById(ctx, payload.ProductId, true)
	if err != nil {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if cart == nil || cart.ID == 0 {
		payload.Price = product.PriceProduct
		_, err := s.CartRepository.SaveCart(ctx, payload)
		if err != nil {
			return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
	} else {
		cart.Quantity += payload.Quantity
		cart.PriceTotal = int64(cart.Quantity) * product.PriceProduct
		_, err := s.CartRepository.UpdateCart(ctx, cart.ID, cart)
		if err != nil {
			return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
	}

	productDto := dto.ProductResponse{
		ID:          product.ID,
		NameProduct: product.NameProduct,
		Category:    product.Category.NameCategory,
		Price:       product.PriceProduct,
	}

	res = &dto.CartResponse{
		ID:       cart.ID,
		Quantity: int64(cart.Quantity),
		Product:  productDto,
		Price:    int64(cart.PriceTotal),
	}

	return res, nil
}

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
