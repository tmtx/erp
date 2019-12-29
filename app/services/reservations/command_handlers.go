package reservations

import (
	"context"

	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/aggregates"
	"github.com/tmtx/res-sys/pkg/bus"
	"github.com/tmtx/res-sys/pkg/event"
	"github.com/tmtx/res-sys/pkg/validator"
)

func (s *Service) Create(p app.CreateReservationParams) (validator.Messages, error) {
	guest := &aggregates.Guest{
		Base:  aggregates.Base{Repository: s.EventRepository},
		Name:  p.GuestName,
		Email: p.GuestEmail,
	}

	if guest.Email != "" {
		err := guest.Restore()
		if err != nil {
			return nil, err
		}
	}

	valid, messages := s.ValidateCreateReservation(guest, p)
	if !valid {
		return messages, nil
	}

	if guest.Id == nil {
		// TODO: api should be something like this:
		// type Guest struct {
		//   Name string `message:"name"`
		//   Email string `message:"email"`
		// }
		// createGuestParams := bus.NewMessageParams(p)
		createGuestParams := bus.MessageParams{
			"name":  p.GuestName,
			"email": p.GuestEmail,
		}
		createGuestCmd := bus.NewCommand(
			app.CreateGuest,
			createGuestParams,
		)

		guestCreationMessages, err := s.CommandBus.DispatchSync(createGuestCmd)
		messages = validator.MergeMessages(messages, guestCreationMessages)

		if err != nil {
			return messages, err
		}
	} else if guest.Name != p.GuestName {
		// TODO: update guest info
	}

	params := bus.MessageParams{
		"guest_name":  p.GuestName,
		"guest_email": p.GuestEmail,
		"start_date":  p.StartDate,
		"end_date":    p.EndDate,
		"space_id":    p.SpaceId,
	}
	event := event.New(app.ReservationCreated, params, event.CreateNewEntityId())

	return nil, s.EventRepository.Store(
		context.Background(),
		event,
	)
}
