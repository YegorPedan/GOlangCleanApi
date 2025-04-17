package server

import (
	"github.com/labstack/echo/v4"
	"net/http"

	eventHttp "github.com/YegorPedan/GOlangCleanApi/internal/event/delivery/http"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	events := v1.Group("/events")

	eventHandler := eventHttp.NewEventHandler()
	eventHttp.MapRoutes(events, eventHandler)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}{
			Message: "ok",
			Status:  http.StatusOK,
		})
	})

	return nil
}
