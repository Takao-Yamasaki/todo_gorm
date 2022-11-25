package main

import (
	"log"
	"net/http"
	"strconv"
	"todo_gorm/domain"
	"todo_gorm/infra"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	Env_load()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

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
			return
		}
		infra.DBCreate(todo)
		c.Redirect(302, "/")
	})
}
