package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetTask(c *gin.Context) {
	var tasks []Task
	var task Task
	task.Title = "Bake some cake"
	task.Body = `- make a dough
	-	Eat everything before baking
	-	Pretend you nerver wanted yo bake something in the first place`

	tasks = append(tasks, task)
	c.JSON(http.StatusOK, gin.H{"task:": tasks})
}

func main() {
	r := gin.Default()
	r.GET("/tasks/", handleGetTask)
	r.Run()
}
