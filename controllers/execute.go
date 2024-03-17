package controllers

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"

	"github.com/NikhilParbat/CC-Compiler-Go/models"
)

func ExecuteCodeHandler(w http.ResponseWriter, r *http.Request) {
	var codeReq models.CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&codeReq); err != nil {
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
	case "go":
		cmd = exec.Command("go", "run")
		cmd.Stdin = strings.NewReader(codeReq.Code)
	case "rb":
		cmd = exec.Command("ruby", "-e", codeReq.Code)
	default:
		http.Error(w, "Unsupported language", http.StatusBadRequest)
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.CodeResponse{
		Output: string(output),
		Error:  "",
	}
	json.NewEncoder(w).Encode(response)
}
