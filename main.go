package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetTask(c *gin.Context) {
	var tasks []Task
	var task Task
	task.Title = "Bake some cake"
	task.Body = "Make a dough"
	tasks = append(tasks, task)
	c.JSON(http.StatusOK, gin.H{"id:": task})
}

func handleCreateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := Create(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id:": id})
}

func main() {
	r := gin.Default()
	r.GET("/tasks/", handleGetTask)
	r.PUT("/tasks/", handleCreateTask)
	r.Run()
}
