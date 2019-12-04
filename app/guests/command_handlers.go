package guests

import (
	"context"
	"fmt"

	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

func (s *Service) Create(p app.CreateGuestParams) (*validator.Messages, error) {
	if isValid, validatorMessages := ValidateCreateGuest(p); !isValid {
		return &validatorMessages, fmt.Errorf("Validation failed")
	}

	ctx := context.Background()
	return nil, s.EventRepository.Store(
		ctx,
		event.NewWithId(app.GuestCreated, p),
	)
}
