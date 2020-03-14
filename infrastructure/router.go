package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/kazu1029/goCleanSample/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())
	eventController := controllers.NewEventController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
	router.GET("/users/:id/events", func(c *gin.Context) { eventController.IndexByUserID(c) })

	router.POST("/events", func(c *gin.Context) { eventController.Create(c) })
	router.GET("/events", func(c *gin.Context) { eventController.Index(c) })
	router.GET("/events/:id", func(c *gin.Context) { eventController.Show(c) })

	Router = router
}
