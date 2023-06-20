package product

import (
	"log"
	"net/http"

	"github.com/ifty123/simple_online_store/internal/factory"
	pkgdto "github.com/ifty123/simple_online_store/pkg/dto"
	"github.com/ifty123/simple_online_store/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: Newservice(f),
	}
}

func (h *handler) GetProducts(c echo.Context) error {
	payload := new(pkgdto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		log.Println("error :", err)
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	res, err := h.service.FindByCategory(c.Request().Context(), payload.Category, &payload.Pagination)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.CustomSuccessBuilder(http.StatusOK, res.Data, "Get product success", &res.PaginationInfo).Send(c)
}
