package router

import (
	"khidr/todo/api/todo"

	"github.com/gin-gonic/gin"
)

func TodoRouteInit(apiGroup *gin.RouterGroup, controller todo.Controller) {
	todos := apiGroup.Group("/todo")
	todos.POST("/", controller.CreateTodo)
	todos.GET("/:id", controller.GetTodoById)
	todos.POST("/query", controller.QueryTodos)
	todos.DELETE("/:id", controller.DeleteTodoById)
	todos.PUT("/:id", controller.UpdateTodo)
}
