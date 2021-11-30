package router

import (
	"khidr/todo/api/todo"

	"github.com/gin-gonic/gin"
)

func TodoRouteInit(apiGroup *gin.RouterGroup, controller todo.Controller) {
	todos := apiGroup.Group("/todo")
	todos.POST("/", controller.CreateTodo)
	todos.GET("/:todo_id", controller.GetTodoById)
	todos.POST("/query", controller.ListTodos)
	todos.DELETE("/:todo_id", controller.DeleteTodoById)
	todos.PUT("/:todo_id", controller.UpdateTodo)
}
