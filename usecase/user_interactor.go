package usecase

import (
	"github.com/go-playground/validator"
	"github.com/kazu1029/goCleanSample/domain"
	"github.com/kazu1029/goCleanSample/interfaces/controllers/request"
)

type UserInteractor struct {
	Repo UserRepository
}

func (interactor *UserInteractor) Add(u *request.User) (err error) {
	user := domain.NewUser(u.NickName, u.Email)
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		return
	}

	_, err = interactor.Repo.Store(user)
	return
}

func (interactor *UserInteractor) Users() (users []domain.User, err error) {
	users, err = interactor.Repo.FindAll()
	return
}

func (interactor *UserInteractor) UserByID(identifier int) (user domain.User, err error) {
	user, err = interactor.Repo.FindByID(identifier)
	return
}
