package reservations

import (
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/aggregates"
	"github.com/tmtx/res-sys/pkg/validator"
)

type existingGuestWithDifferentData struct {
	Params app.CreateReservationParams
	*aggregates.Guest
}

func (v existingGuestWithDifferentData) Validate() (bool, validator.Message) {
	// shouldn't happen as guest aggregate is restored by email
	if v.Guest.Email != v.Params.GuestEmail {
		return false, validator.Message("Existing guest with different email")
	}

	if v.Guest.Name != v.Params.GuestName {
		return false, validator.Message("Existing guest with different name")
	}

	return true, ""
}

func (s *Service) ValidateCreateReservation(
	guest *aggregates.Guest,
	p app.CreateReservationParams,
) (bool, validator.Messages) {
	emailValidators := app.EmailValidators(p.GuestEmail)
	nameValidators := app.NameValidators(p.GuestName)

	emailValidators = append(
		emailValidators,
		existingGuestWithDifferentData{Params: p, Guest: guest},
	)

	return validator.ValidateGroup(validator.Group{
		"name":  nameValidators,
		"email": emailValidators,
	})
}
