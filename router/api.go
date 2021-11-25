package router

import (
	"khidr/todo/api/todo"
	"khidr/todo/persistence"
	"log"

	"github.com/gin-gonic/gin"
)

func InitAPI(dsn string) *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	postgresConfig, err := persistence.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}
	todoSvc := todo.NewService(postgresConfig)
	todoHandler := todo.NewHandler()

	todoController := todo.NewController(todoSvc, todoHandler)
	todoRouter := apiV1.Group("/todo")
	todoRouter.POST("/", todoController.CreateTodo)

	return router
}
