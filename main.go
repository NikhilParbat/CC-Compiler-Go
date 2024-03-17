package main

import (
	"fmt"
	"net/http"

	"github.com/NikhilParbat/CC-Compiler-Go/controllers"
)

func main() {
	// Define your route handlers
	http.HandleFunc("/execute", controllers.ExecuteCodeHandler)

	// Start the server with the CORS middleware
	fmt.Println("Server listening on port 5000...")
	http.ListenAndServe(":5000", nil)
}
func Handler() {
	main()
}
