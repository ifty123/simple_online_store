package cart

import (
	"context"
	"log"

	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/factory"
	"github.com/ifty123/simple_online_store/internal/repository"
	pkgdto "github.com/ifty123/simple_online_store/pkg/dto"

	"github.com/ifty123/simple_online_store/pkg/util/response"
)

type Service struct {
	CartRepository    repository.CartRepository
	ProductRepository repository.ProductRepository
}

type CartService interface {
	FindCartByUserId(ctx context.Context, userId uint) ([]*dto.CartResponse, int64, error)
	SaveCart(ctx context.Context, payload *dto.Cart) (*dto.CartResponse, error)
	DeleteProductById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.CartDeleteResponse, error)
}

func Newservice(f *factory.Factory) CartService {
	return &Service{
		CartRepository:    f.CartRepository,
		ProductRepository: f.ProductRepository,
	}
}

func (s *Service) SaveCart(ctx context.Context, payload *dto.Cart) (*dto.CartResponse, error) {

	var res *dto.CartResponse

	//cek payload quantity, jika < 0, maka default ke 1
	if payload.Quantity <= 0 {
		payload.Quantity = 1
	}

	cart, err := s.CartRepository.FindByProductId(ctx, payload.ProductId, payload.UserId)
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
		saveCart, err := s.CartRepository.SaveCart(ctx, payload)
		if err != nil {
			return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}

		//masukkan ke cart untuk response
		cart.ID = saveCart.ID
		cart.PriceTotal = saveCart.PriceTotal
		cart.Quantity = saveCart.Quantity

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

func (s *Service) FindCartByUserId(ctx context.Context, userId uint) ([]*dto.CartResponse, int64, error) {

	var totalCart int64

	cart, err := s.CartRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, 0, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
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

			totalCart += i.PriceTotal
		}
	}

	return productRes, totalCart, nil
}

func (s *Service) DeleteProductById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.CartDeleteResponse, error) {

	//find by id
	cart, err := s.CartRepository.FindById(ctx, payload.ID)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	_, err = s.CartRepository.Destroy(ctx, cart)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result := &dto.CartDeleteResponse{
		ID:        cart.ID,
		ProductId: cart.ProductId,
		Deleted:   cart.DeletedAt.Time,
	}

	return result, nil
}
