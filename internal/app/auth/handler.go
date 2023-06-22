package auth

import (
	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/factory"
	"github.com/ifty123/simple_online_store/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	Service AuthService
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		Service: NewService(f),
	}
}

// @Tags Auth
// @Summary API Login By Email And Password
// @Router /auth/login [post]
// @Param request body dto.EmailAndPasswordReq true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Success 200 {object} dto.UserWithJWTResponse
// @Failure 400 {object} response.Error
func (h *handler) LoginByEmailAndPassword(c echo.Context) error {
	payload := new(dto.EmailAndPasswordReq)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	employee, err := h.Service.LoginByEmailAndPassword(c.Request().Context(), payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(employee).Send(c)
}

// @Tags Auth
// @Summary API Register By Email And Password
// @Router /auth/signup [post]
// @Param request body dto.RegisterUserReq true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} response.Error
func (h *handler) RegisterByEmailAndPassword(c echo.Context) error {
	payload := new(dto.RegisterUserReq)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	employee, err := h.Service.RegisterByEmailAndPassword(c.Request().Context(), payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(employee).Send(c)
}
