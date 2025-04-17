package event

import "github.com/YegorPedan/GOlangCleanApi/internal/model"

type EventUC interface {
	FindById(id int64) (*model.Event, error)
}
