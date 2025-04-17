package http

import (
	"github.com/YegorPedan/GOlangCleanApi/internal/event"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type EventHandler struct {
	eventUC event.EventUC
}

func NewEventHandler() event.Handler {
	return &EventHandler{}
}

func (e *EventHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
				Status  int    `json:"status"`
			}{
				Message: "Bad Request",
				Status:  http.StatusBadRequest,
			})
		}

		event1, err := e.eventUC.FindById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
				Status  int    `json:"status"`
			}{
				Message: "Not Found",
				Status:  http.StatusNotFound,
			})
		}

		return c.JSON(http.StatusOK, event1)
	}
}
