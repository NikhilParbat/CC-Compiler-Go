package handler

import (
	"fmt"
	"net/http"

	"github.com/NikhilParbat/CC-Compiler-Go/controllers"
)

// handler function to handle HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	// Define your route handlers
	switch r.URL.Path {
	case "/execute":
		fmt.Fprintln(w, "Hello from Execute!")
	default:
		fmt.Fprintln(w, "Hello from Default!")
		controllers.ExecuteCodeHandler(w, r)
	}
}

// CORS middleware function to set necessary headers
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set headers to allow cross-origin requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Continue with the next handler
		next(w, r)
	}
}

func main() {
	// Start the serverless function
	http.HandleFunc("/", CORS(Handler))
}
