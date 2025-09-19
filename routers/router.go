package routers

import (
	"buble-master/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//Api分组 版本v1
	v1Group := r.Group("v1")
	{
		//代办事项路由
		v1Group.POST("/todo", handlers.CreateTodo)
		v1Group.DELETE("/todo/:id", handlers.DeleteTodo)
		v1Group.PUT("/todo/:id", handlers.UpdateTodoStatus)
		v1Group.GET("/todo", handlers.GetAllTodos)
	}
	return r
}
