package product

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetProducts)
	g.GET("/category", h.GetCategory)
}
