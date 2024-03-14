package main

import (
	"fmt"

	"github.com/NikhilParbat/CC-Compiler-Go/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://example.com"}           // Add allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"} // Add allowed HTTP methods
	config.AllowHeaders = []string{"Origin"}                       // Add allowed headers

	router.Use(cors.New(config))

	router.POST("/execute", controllers.ExecuteCode)

	fmt.Println("Server listening on port 5000...")
	router.Run(":5000")
}
