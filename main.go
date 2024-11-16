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


	r.Run(":8080")
}

