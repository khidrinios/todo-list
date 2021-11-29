package router

import (
	"khidr/todo/api/item"
	"khidr/todo/api/todo"
	"khidr/todo/persistence"
	"log"

	"github.com/gin-gonic/gin"
)

func InitAPI(dsn string) *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	//init persistence
	postgresConfig, err := persistence.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}

	//init services
	todoSvc := todo.NewService(postgresConfig)
	itemSvc := item.NewService(postgresConfig)

	// init handlers
	todoHandler := todo.NewHandler()
	itemHandler := item.NewHandler()

	//init controllers
	todoController := todo.NewController(todoSvc, todoHandler)
	itemController := item.NewController(itemSvc, itemHandler)

	//init routers
	TodoRouteInit(apiV1, todoController)
	ItemRouteInit(apiV1, itemController)

	return router
}
