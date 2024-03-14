package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type CodeRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type CodeResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}

func executeCode(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var codeReq CodeRequest
	err := decoder.Decode(&codeReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cmd *exec.Cmd
	switch codeReq.Language {
	case "js":
		cmd = exec.Command("node", "-e", codeReq.Code)
	case "py":
		cmd = exec.Command("python", "-")
		cmd.Stdin = strings.NewReader(codeReq.Code)
	case "c":
		cmd = exec.Command("gcc", "-o", "temp", "-x", "c", "-")
		cmd.Stdin = strings.NewReader(codeReq.Code)
		err := cmd.Run()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cmd = exec.Command("./temp")
	default:
		http.Error(w, "Unsupported language", http.StatusBadRequest)
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := CodeResponse{
		Output: string(output),
		Error:  "",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	// Remove the temporary executable file if it exists
	if _, err := os.Stat("./temp"); err == nil {
		os.Remove("./temp")
	}
}

func main() {
	http.HandleFunc("/execute", executeCode)
	fmt.Println("Server listening on port 5000...")
	http.ListenAndServe(":5000", nil)
}
