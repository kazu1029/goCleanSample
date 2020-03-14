package usecase

import (
	"github.com/go-playground/validator"
	"github.com/kazu1029/goCleanSample/domain"
	"github.com/kazu1029/goCleanSample/interfaces/controllers/request"
)

type EventInteractor struct {
	Repo EventRepository
}

func (interactor *EventInteractor) Add(e *request.Event) (err error) {
	event := domain.NewEvent(e.UserID, e.Title, e.Description)
	validate := validator.New()
	err = validate.Struct(event)
	if err != nil {
		return
	}

	_, err = interactor.Repo.Store(event)
	return
}

func (interactor *EventInteractor) Events() (events []domain.Event, err error) {
	events, err = interactor.Repo.FindAll()
	return
}

func (interactor *EventInteractor) EventByID(identifier int) (event domain.Event, err error) {
	event, err = interactor.Repo.FindByID(identifier)
	return
}

func (interactor *EventInteractor) EventsByUserID(user_id int) (events []domain.Event, err error) {
	events, err = interactor.Repo.FindAllByUserID(user_id)
	return
}
