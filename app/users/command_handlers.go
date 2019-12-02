package users

import (
	"context"
	"fmt"

	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/pkg/event"
	"github.com/tmtx/erp/pkg/validator"
)

func (s *Service) Login(p app.LoginUserParams) (error, *validator.Messages) {
	if isValid, validatorMessages := ValidateLoginUser(p); !isValid {
		fmt.Println(validatorMessages)
		return fmt.Errorf("Validation failed"), validatorMessages
	}

	return s.EventRepository.Store(
		context.Background(),
		event.New(app.GuestCreated, p),
	), nil
}
