package router

import (
	"khidr/todo/api/item"
	"khidr/todo/api/todo"
	"khidr/todo/db"
	"log"

	"github.com/gin-gonic/gin"
)

func InitAPI(dsn string) *gin.Engine {
	router := gin.Default()
	HealthCheckInitRoute(router)
	apiV1 := router.Group("/api/v1")

	//init persistence
	database, err := db.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}

	//init repos
	todoRepo := todo.NewRepository(database)
	itemRepo := item.NewRepository(database)

	//init services
	todoSvc := todo.NewService(todoRepo)
	itemSvc := item.NewService(itemRepo)

	// init handlers
	todoHandler := todo.NewHandler()
	itemHandler := item.NewHandler()

	//init controllers
	todoController := todo.NewController(todoSvc, itemSvc, todoHandler)
	itemController := item.NewController(itemSvc, itemHandler)

	//init routers
	TodoRouteInit(apiV1, todoController)
	ItemRouteInit(apiV1, itemController)

	return router
}
