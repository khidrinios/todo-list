package router

import (
	"khidr/todo/api/item"

	"github.com/gin-gonic/gin"
)

func ItemRouteInit(apiGroup *gin.RouterGroup, controller item.Controller) {
	items := apiGroup.Group("todo/:todo_id/item")
	items.POST("/", controller.AddToTodo)
	items.GET("/:item_id", controller.Get)
	items.DELETE("/:item_id", controller.Delete)
}
