package http

import "github.com/labstack/echo/v4"

func MapRoutes(g *echo.Group, h *EventHandler) {
	g.GET("/:id", h.Get())
}
