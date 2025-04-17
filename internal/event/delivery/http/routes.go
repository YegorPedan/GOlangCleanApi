package http

import (
	"github.com/YegorPedan/GOlangCleanApi/internal/event"
	"github.com/labstack/echo/v4"
)

func MapRoutes(g *echo.Group, h event.Handler) {
	g.GET("/:id", h.Get())
}
