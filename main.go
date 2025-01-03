package main

import (
	"devjudge/go-in-docker/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	r.GET("/api/leads/:lead_id", handler.GetLeadById)
	r.POST("/api/leads/", handler.CreateLead)
	r.Run(":8080")
}
