package usecase

import "github.com/kazu1029/goCleanSample/domain"

type EventRepository interface {
	Store(*domain.Event) (int, error)
	FindByID(int) (domain.Event, error)
	FindAll() ([]domain.Event, error)
	FindAllByUserID(int) ([]domain.Event, error)
}
