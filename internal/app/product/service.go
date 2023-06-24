package product

import (
	"context"

	dto_internal "github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/factory"
	"github.com/ifty123/simple_online_store/internal/repository"

	pkgdto "github.com/ifty123/simple_online_store/pkg/dto"
	"github.com/ifty123/simple_online_store/pkg/util/response"
)

type Service struct {
	ProductRepository  repository.ProductRepository
	CategoryRepository repository.CategoryRepository
}

type ProductService interface {
	FindByCategory(ctx context.Context, category uint, pagination *pkgdto.Pagination) (*pkgdto.SearchGetResponse[dto_internal.ProductResponse], error)
	GetAllCategory(ctx context.Context) ([]*dto_internal.CategoryResponse, error)
}

func Newservice(f *factory.Factory) Service {
	return Service{
		ProductRepository:  f.ProductRepository,
		CategoryRepository: f.CategoryRepository,
	}
}

func (s *Service) FindByCategory(ctx context.Context, category uint, pagination *pkgdto.Pagination) (*pkgdto.SearchGetResponse[dto_internal.ProductResponse], error) {

	product, paginate, err := s.ProductRepository.FindByCategory(ctx, category, pagination)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	productRes := []dto_internal.ProductResponse{}
	for _, i := range product {
		productRes = append(productRes, dto_internal.ProductResponse{
			ID:          i.ID,
			NameProduct: i.NameProduct,
			Price:       i.PriceProduct,
			Category:    i.Category.NameCategory,
		})
	}

	res := new(pkgdto.SearchGetResponse[dto_internal.ProductResponse])

	res.Data = productRes
	res.PaginationInfo = *paginate

	return res, nil
}

func (s *Service) GetAllCategory(ctx context.Context) ([]*dto_internal.CategoryResponse, error) {

	category, err := s.CategoryRepository.GetAll(ctx)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	productRes := []*dto_internal.CategoryResponse{}
	for _, i := range category {
		productRes = append(productRes, &dto_internal.CategoryResponse{
			ID:           i.ID,
			NameCategory: i.NameCategory,
		})
	}

	return productRes, nil
}
