package API

import (
	"encoding/json"
	"net/http"
)

func BuildJsonResponse(status bool, message string, content interface{}, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response{
		Status:  status,
		Message: message,
		Content: content,
	})
}
