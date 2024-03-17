package handler

import (
	"net/http"

	"github.com/NikhilParbat/CC-Compiler-Go/controllers"
)

// Handler function to handle HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set headers to allow cross-origin requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Execute the actual logic of the Handler
	controllers.ExecuteCodeHandler(w, r)
}

func init() {
	// Start the serverless function
	http.HandleFunc("/execute", Handler)
	http.ListenAndServe(":8080", nil)
}
