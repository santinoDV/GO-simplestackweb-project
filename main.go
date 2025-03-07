package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")

	e := gin.Default()

	InitDatabase()

	// Load HTML templates from the "templates" folder
	e.LoadHTMLGlob("templates/*")

	e.GET("/", func(c *gin.Context) {
		log.Println("logging HTML")
		todos := ReadToDoList()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	e.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		status := c.PostForm("status")
		id, err := CreateToDo(title, status)
		if err != nil {
			log.Printf("Error creating todo: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create todo"})
			return
		}

		c.HTML(http.StatusOK, "task.html", gin.H{
			"title":  title,
			"status": status,
			"id":     id,
		})
	})

	e.DELETE("/todos/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, err := strconv.ParseInt(param, 10, 64)

		if err != nil {
			log.Printf("Error parsing ID: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		DeleteToDo(id)
		c.Status(http.StatusNoContent) // Respondo with 204
	})

	e.Run(":8080")
}

/**/
