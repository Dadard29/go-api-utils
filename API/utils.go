package API

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	MethodNotAllowed    = "method not allowed"
	MissingParameter    = "missing parameter"
	InvalidHeaderBearer = "invalid header (authorization bearer)"
	InvalidHeaderBasic  = "invalid header (authorization basic)"

	jsonIndent = "    "
)

func (a API) BuildJsonResponse(status bool, message string, content interface{}, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", a.Service.CorsOrigin())
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", jsonIndent)
	return encoder.Encode(response{
		Status:  status,
		Message: message,
		Content: content,
	})
}

func (a API) BuildErrorResponse(httpCode int, message string, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", a.Service.CorsOrigin())
	w.WriteHeader(httpCode)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", jsonIndent)
	return encoder.Encode(response{
		Status:  false,
		Message: message,
		Content: struct{}{},
	})
}

func (a API) BuildMissingParameter(w http.ResponseWriter) error {
	return a.BuildErrorResponse(http.StatusBadRequest, MissingParameter, w)
}

func (a API) BuildMethodNotAllowedResponse(w http.ResponseWriter) error {
	return a.BuildErrorResponse(http.StatusMethodNotAllowed, MethodNotAllowed, w)
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
