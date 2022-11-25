package main

import (
	"net/http"
	"strconv"
	"todo_gorm/infra"
	"todo_gorm/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	
	db := infra.DBInit()
	defer db.Close()
	
	router.GET("/", func(c *gin.Context) {
		todos := infra.DBRead()
		c.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})
	
	//POSTの処理
	router.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")
		rawStatus := c.PostForm("status")
		id, err := strconv.Atoi(rawStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		status := domain.Status(id)
		rawTime := c.PostForm("deadline")
		deadline, err := strconv.Atoi(rawTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		todo := domain.Todo{Text: text, Status: status, Deadline: deadline}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		infra.DBCreate(todo)
		c.Redirect(302, "/")
	})
}
