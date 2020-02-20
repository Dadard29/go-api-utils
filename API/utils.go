package API

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	MethodNotAllowed    = "method not allowed"
	InvalidHeaderBearer = "invalid header (authorization bearer)"
	InvalidHeaderBasic  = "invalid header (authorization basic)"
)

func BuildJsonResponse(status bool, message string, content interface{}, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response{
		Status:  status,
		Message: message,
		Content: content,
	})
}

func BuildErrorResponse(httpCode int, message string, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	return json.NewEncoder(w).Encode(response{
		Status:  false,
		Message: message,
		Content: struct{}{},
	})
}

func BuildMethodNotAllowedResponse(w http.ResponseWriter) {
	_ = BuildErrorResponse(http.StatusMethodNotAllowed, MethodNotAllowed, w)
}

func CheckHttpMethod(r *http.Request, expectedMethod string) bool {
	if r.Method != expectedMethod {
		return false
	}
	return true
}

func ParseJsonBody(r *http.Request, object interface{}) error {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, object)
}
