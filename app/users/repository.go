package users

import (
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/event"
)

type repo struct {
	eventRepository event.Repository
}

func NewRepository(eventRepository event.Repository) app.AggregateRootRepository {
	return &repo{eventRepository}
}

func (r *repo) RestoreAggregateRootById(id app.UUID) (error, *app.Aggregate) {
	return nil, nil
}

func (r *repo) RestoreAggregateRootByCredentials(id app.UUID) (error, *app.Aggregate) {
	return nil, nil
}
