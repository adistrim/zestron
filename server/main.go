package main

import (
	"net/http"
	"os"
	"strings"
	"zestron-server/handlers"
	"zestron-server/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	config := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(config))

	chatManager := services.NewChatManager()

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

	router.GET("/origins", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"allowedOrigins": allowedOrigins,
			"yourOrigin":     c.Request.Header.Get("Origin"),
		})
	})

	router.GET("/ws/generate", handlers.WebSocketHandler(chatManager, allowedOrigins))

	router.Run(":8080")
}
