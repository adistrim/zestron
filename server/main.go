package main

import (
	"net/http"
	"zestron-server/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Static("/public", "./public")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/404")
	})

	router.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	router.GET("/404", func(c *gin.Context) {
		c.File("./public/404.html")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/generate", handlers.GenerateHandler)

	router.Run(":8080")
}
