package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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

	r.Run(":8080")
}

