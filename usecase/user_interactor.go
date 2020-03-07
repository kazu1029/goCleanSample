package usecase

import "github.com/kazu1029/goCleanSample/domain"

type UserInteractor struct {
	Repo UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.Repo.Store(u)
	return
}

func (interactor *UserInteractor) Users() (users []domain.User, err error) {
	users, err = interactor.Repo.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.Repo.FindById(identifier)
	return
}
