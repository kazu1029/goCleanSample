package usecase

import "github.com/kazu1029/goCleanSample/domain"

type UserRepository interface {
	Store(*domain.User) (int, error)
	FindByID(int) (domain.User, error)
	FindAll() ([]domain.User, error)
}
