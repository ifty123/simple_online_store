package cart

import (
	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/factory"
	"github.com/ifty123/simple_online_store/pkg/util"
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

func (h *handler) GetCart(c echo.Context) error {

	//get auth : userId
	authHeader := c.Request().Header.Get("Authorization")

	jwtClaims, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Unauthorized, err).Send(c)
	}

	res, err := h.service.FindCartByUserId(c.Request().Context(), jwtClaims.UserID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(res).Send(c)
}

func (h *handler) SaveCart(c echo.Context) error {

	payload := new(dto.Cart)

	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	//get auth : userId
	authHeader := c.Request().Header.Get("Authorization")

	jwtClaims, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload.UserId = jwtClaims.UserID

	res, err := h.service.SaveCart(c.Request().Context(), payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(res).Send(c)
}
