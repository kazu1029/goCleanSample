package usecase

import "github.com/kazu1029/goCleanSample/domain"

type UserInteractor struct {
	repo UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.repo.Store(u)
	return
}

func (interactor *UserInteractor) Users() (user []domain.User, err error) {
	user, err = interactor.repo.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.repo.FindById(identifier)
	return
}
