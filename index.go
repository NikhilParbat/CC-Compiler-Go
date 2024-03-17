package handler

import (
	"fmt"
	"net/http"

	"github.com/NikhilParbat/CC-Compiler-Go/controllers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Define your route handlers
	switch r.URL.Path {
	case "/execute":
		fmt.Fprintln(w, "Hello!")
	default:
		controllers.ExecuteCodeHandler(w, r)
	}
}

func main() {
	// Start the serverless function
	http.HandleFunc("/", Handler)
}
