package guests

import (
	"context"
	"fmt"

	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/pkg/bus"
	"github.com/tmtx/res-sys/pkg/event"
	"github.com/tmtx/res-sys/pkg/validator"
)

func (s *Service) Create(p app.CreateGuestParams) (validator.Messages, error) {
	if isValid, validatorMessages := ValidateCreateGuest(p); !isValid {
		return validatorMessages, fmt.Errorf("Validation failed")
	}

	ctx := context.Background()

	params := bus.MessageParams{
		"name":  p.Name,
		"email": p.Email,
	}
	event := event.New(app.GuestCreated, params, event.CreateNewEntityId())

	return nil, s.EventRepository.Store(
		ctx,
		event,
	)
}
