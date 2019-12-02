package guests

import (
	"context"
	"fmt"

	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

func (s *Service) Create(p app.CreateGuestParams) (error, *validator.Messages) {
	if isValid, validatorMessages := ValidateCreateGuest(p); !isValid {
		return fmt.Errorf("Validation failed"), &validatorMessages
	}

	p.Id = app.NewUUID()

	ctx := context.Background()
	return s.EventRepository.Store(
		ctx,
		event.New(app.GuestCreated, p),
	), nil
}
