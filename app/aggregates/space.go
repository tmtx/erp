package aggregates

import (
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

type Space struct {
	Base
	Id int `json:"id"`
}

func (ag *Space) GetTargetEvents() []bus.MessageKey {
	return []bus.MessageKey{}
}

func (ag *Space) ApplyEvent(e event.Event) {
	// TODO: implement
}

func (ag *Space) Validate() (bool, *validator.Messages) {
	// TODO: implement
	return false, nil
}

func (ag *Space) CanBeRestored() bool {
	// TODO: implement
	return true
}

func (ag *Space) Restore() error {
	// TODO: implement
	return nil
}
