package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/goCleanSample/interfaces/controllers/request"
	"github.com/kazu1029/goCleanSample/interfaces/iRepository"
	"github.com/kazu1029/goCleanSample/usecase"
)

type EventController struct {
	Interactor usecase.EventInteractor
}

func NewEventController(sqlHandler iRepository.SqlHandler) *EventController {
	return &EventController{
		Interactor: usecase.EventInteractor{
			Repo: &iRepository.EventRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *EventController) Create(c Context) {
	e := request.Event{}
	c.Bind(&e)

	err := controller.Interactor.Add(&e)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, gin.H{"message": "created"})
}

func (controller *EventController) Index(c Context) {
	events, err := controller.Interactor.Events()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, events)
}

func (controller *EventController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	event, err := controller.Interactor.EventById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, event)
}
