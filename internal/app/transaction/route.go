package transaction

import (
	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/middleware"
	"github.com/ifty123/simple_online_store/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {

	g.Use(middleware.JWTMiddleware(dto.JWTClaims{}, util.JWT_SECRET))

	g.POST("/add", h.SaveTransaction)
	g.GET("", h.GetTransactionByUserId)
}
