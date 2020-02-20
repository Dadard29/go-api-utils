package test

import (
	"fmt"
	"github.com/Dadard29/go-api-utils/auth"
	"testing"
	"time"
)

type Payload struct {
	Infos string
}

var pl = Payload{
	Infos: "infos",
}

var secret = "secret"
var issuer = "issuer"
var subject = "subject"
var audience = []string{"audience"}
var validity = 24 * time.Hour

var path = "auth/test/private.pem"

func TestJwtPlain(t *testing.T) {
	jwt, err := auth.NewJwtHS256(
		secret,
		issuer,
		subject,
		audience,
		validity,
		pl)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(jwt))

	newPl, err := auth.VerifyJwtHS256(jwt, secret)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*newPl)
}

func TestJwtCiphered(t *testing.T) {
	jwt, err := auth.NewJwtHS256(
		secret,
		issuer,
		subject,
		audience,
		validity,
		pl)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(jwt))

	cipheredToken, err := auth.CipherJwtWithJwe(path, jwt)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(cipheredToken))
	deciphered, err := auth.DecipherJwtWithJwe(path, cipheredToken)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(deciphered))
}
