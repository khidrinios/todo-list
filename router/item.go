package router

import (
	"khidr/todo/api/item"

	"github.com/gin-gonic/gin"
)

func ItemRouteInit(apiGroup *gin.RouterGroup, controller item.Controller) {
	items := apiGroup.Group("todo/:todo_id/item")
	items.POST("/", controller.AddItemToTodo)
	items.GET("/:item_id", controller.GetItemByIdAndTodoId)
}
