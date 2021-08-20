package controllers

import (
	"net/http"

	"github.com/DanBillson/godo/models"
	"github.com/gin-gonic/gin"
)

type CreateTodoInput struct {
	Title string `json:"title" binding:"required"`
}

// GET /todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	models.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// POST /todo
func CreateTodo(c *gin.Context) {
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Title: input.Title, Done: false}
	models.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// GET /todo/:id
func GetTodo(c *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// PATCH /todo/:id
func ToggleTodo(c *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
	}

	models.DB.Model(&todo).Update("Done", !todo.Done)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// DELETE /todo/:id
func DeleteTodo(c *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
	}

	models.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
