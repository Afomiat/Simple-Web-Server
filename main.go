package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.GET("/name", func(c *gin.Context) {
		c.String(http.StatusOK, "Afomia Tadesse")
	})

	r.GET("/hobby", func(c *gin.Context) {
		c.String(http.StatusOK, "I love exploring new places!")
	})

	r.GET("/dream", func(c *gin.Context) {
		c.String(http.StatusOK, "I dream of becoming a Backend Developer!")
	})

	r.Run(":" + port)
}
