package auth

import (
	"encoding/base64"
	"errors"
	"github.com/Dadard29/go-api-utils/API"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"net/http"
	"strings"
)

var logger = log.NewLogger("AUTH", logLevel.DEBUG)

func ParseBasicAuth(r *http.Request) (username string, password string, err error) {
	header := r.Header.Get("Authorization")

	splitted := strings.Split(header, " ")
	if len(splitted) < 2 {
		return "", "", errors.New(API.InvalidHeaderBasic)
	}

	b64value := splitted[1]

	decodedBytes, err := base64.StdEncoding.DecodeString(b64value)
	logger.CheckErr(err)

	decoded := string(decodedBytes)
	userAndPass := strings.Split(decoded, ":")

	return userAndPass[0], userAndPass[1], nil
}

func ParseBearerToken(r *http.Request) (token string, err error) {
	header := r.Header.Get("Authorization")

	splitted := strings.Split(header, " ")
	if len(splitted) < 2 {
		return "", errors.New(API.InvalidHeaderBearer)
	}

	return splitted[1], nil
}

func ParseApiKey(r *http.Request, key string, inHeader bool) (token string) {
	// if inHeader true, search the key in the headers
	// else, search in the query params
	if inHeader {
		token = r.Header.Get(key)
	} else {
		token = r.URL.Query().Get(key)
	}
	return
}
