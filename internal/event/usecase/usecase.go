package usecase

import (
	"github.com/YegorPedan/GOlangCleanApi/internal/event"
	"github.com/YegorPedan/GOlangCleanApi/internal/model"
)

type EventUCase struct {
	eventUC event.EventUC
}

func NewEventUCase(uc event.EventUC) event.EventUC {
	return &EventUCase{eventUC: uc}
}

func (u *EventUCase) FindById(id int64) (*model.Event, error) {
	return &model.Event{
		ID:    id,
		Title: "Nice",
	}, nil
}
