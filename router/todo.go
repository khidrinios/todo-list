package router

import (
	"khidr/todo/api/todo"

	"github.com/gin-gonic/gin"
)

func TodoRouteInit(apiGroup *gin.RouterGroup, controller todo.Controller) {
	todos := apiGroup.Group("/todo")
	todos.POST("/", controller.Create)
	todos.GET("/:todo_id", controller.GetById)
	todos.POST("/query", controller.List)
	todos.DELETE("/:todo_id", controller.DeleteById)
	todos.PUT("/:todo_id", controller.Update)
}
