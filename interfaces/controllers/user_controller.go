package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/goCleanSample/interfaces/controllers/request"
	"github.com/kazu1029/goCleanSample/interfaces/iRepository"
	"github.com/kazu1029/goCleanSample/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler iRepository.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			Repo: &iRepository.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := request.User{}
	c.Bind(&u)

	err := controller.Interactor.Add(&u)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, gin.H{"message": "created"})
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserByID(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, user)
}
