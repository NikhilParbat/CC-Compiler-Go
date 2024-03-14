package main

import (
	"fmt"
	"net/http"

	"github.com/NikhilParbat/CC-Compiler-Go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	router.POST("/execute", controllers.ExecuteCode)

	fmt.Println("Server listening on port 5000...")
	router.Run(":5000")
}
