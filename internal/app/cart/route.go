package cart

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetCart)
	g.PUT("/add", h.SaveCart)
}
