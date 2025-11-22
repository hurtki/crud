package tasksHandler

import (
	"encoding/json"
	"net/http"
)

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]any{
		"error": msg,
	}

	json.NewEncoder(w).Encode(resp)
}
