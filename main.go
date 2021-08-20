package main

import (
	"github.com/DanBillson/godo/controllers"
	"github.com/DanBillson/godo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	r.GET("/todos", controllers.GetTodos)
	r.POST("/todo", controllers.CreateTodo)
	r.GET("/todo/:id", controllers.GetTodo)
	r.PATCH("/todo/:id", controllers.ToggleTodo)
	r.DELETE("/todo/:id", controllers.DeleteTodo)

	r.Run()
}
