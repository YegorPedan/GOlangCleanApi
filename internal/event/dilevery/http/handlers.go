package http

import (
	"github.com/YegorPedan/GOlangCleanApi/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type EventHandler struct{}

func NewEventHandler() *EventHandler {
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

		event := &model.Event{
			ID:    id,
			Title: "EXample",
		}
		return c.JSON(http.StatusOK, event)
	}
}
