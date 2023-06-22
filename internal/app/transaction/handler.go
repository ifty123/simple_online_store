package transaction

import (
	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/factory"
	pkgdto "github.com/ifty123/simple_online_store/pkg/dto"
	"github.com/ifty123/simple_online_store/pkg/util"
	"github.com/ifty123/simple_online_store/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service TransactionService
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: Newservice(f),
	}
}

// @Tags Transaction
// @Summary API Save Transaction
// @Router /transaction/add [post]
// @Param request body dto.TransactionReq true "Payload Body [RAW]"
// @Security SH256
// @Accept json
// @Produces json
// @Success 200 {object} dto.TransactionResponse
// @Failure 400 {object} response.Error
func (h *handler) SaveTransaction(c echo.Context) error {

	payload := new(dto.TransactionReq)

	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	//get auth : userId
	authHeader := c.Request().Header.Get("Authorization")

	jwtClaims, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload.UserId = jwtClaims.UserID

	res, err := h.service.SaveTransaction(c.Request().Context(), payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(res).Send(c)
}

// @Tags Transaction
// @Summary API Get Transaction
// @Router /transaction [get]
// @Security SH256
// @Accept json
// @Produces json
// @Success 200 {object} []dto.TransactionResponse
// @Failure 400 {object} response.Error
func (h *handler) GetTransactionByUserId(c echo.Context) error {

	//get auth : userId
	authHeader := c.Request().Header.Get("Authorization")

	jwtClaims, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Unauthorized, err).Send(c)
	}

	res, err := h.service.FindTransactionByUserId(c.Request().Context(), jwtClaims.UserID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(res).Send(c)
}

// @Tags Transaction
// @Summary API Update Transaction
// @Router /transaction/update/:id [put]
// @Param id path int true "transaction_id"
// @Security SH256
// @Accept json
// @Produces json
// @Success 200 {object} dto.TransactionResponse
// @Failure 400 {object} response.Error
func (h *handler) UpdateTransactionById(c echo.Context) error {

	payload := new(pkgdto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	//get auth : userId
	authHeader := c.Request().Header.Get("Authorization")

	jwtClaims, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Unauthorized, err).Send(c)
	}

	res, err := h.service.UpdateTransactionById(c.Request().Context(), jwtClaims.UserID, payload.ID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(res).Send(c)
}
