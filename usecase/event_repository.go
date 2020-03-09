package usecase

import "github.com/kazu1029/goCleanSample/domain"

type EventRepository interface {
	Store(*domain.Event) (int, error)
	FindById(int) (domain.Event, error)
	FindAll() ([]domain.Event, error)
}
