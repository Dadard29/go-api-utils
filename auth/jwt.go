package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/gbrlsnchs/jwt"
	"github.com/square/go-jose"
	"io/ioutil"
	"time"
)

type JwtPayload struct {
	jwt.Payload
	Infos interface{}
}

func NewJwtHS256(secret string,
	issuer string, subject string, audience []string, validityDuration time.Duration,
	payload interface{}) ([]byte, error) {

	secretHmac := jwt.NewHS256([]byte(secret))
	now := time.Now()

	pl := struct {
		jwt.Payload
		Infos interface{}

	} {
		Payload: jwt.Payload{
			Issuer:         issuer,
			Subject:        subject,
			Audience:       audience,
			ExpirationTime: jwt.NumericDate(now.Add(validityDuration)),
			NotBefore:      nil,
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "",
		},
		Infos: payload,
	}

	token, err := jwt.Sign(pl, secretHmac)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func VerifyJwtHS256(token []byte, secret string) (*JwtPayload, error) {
	secretHmac := jwt.NewHS256([]byte(secret))
	var pl JwtPayload

	expValidator := jwt.ExpirationTimeValidator(time.Now())
	validator := jwt.ValidatePayload(&pl.Payload, expValidator)

	_, err := jwt.Verify(token, secretHmac, &pl, jwt.ValidateHeader, validator)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}


// private key generated with
// openssl genrsa 2048 | openssl pkcs8 -topk8 -nocrypt -out private.pem
func readPrivateKeyFile(pathPrivateKeyFile string) (*rsa.PrivateKey, error) {
	priv, err := ioutil.ReadFile(pathPrivateKeyFile)
	if err != nil {
		return nil, errors.New("error while reading private key file")
	}
	privPem, _ := pem.Decode(priv)
	if privPem.Type != "PRIVATE KEY" {
		return nil, errors.New("the provided file is not a RSA private key format")
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(privPem.Bytes)
	if err != nil {
		return nil, err
	}

	privKey, check := parsedKey.(*rsa.PrivateKey)
	if ! check {
		return nil, errors.New("unable to parse private key")
	}

	return privKey, nil
}

// return base64-encoded a ciphered JWT with PKCS8 with a RSA private key
func CipherJwtWithJwe(pathPrivateKeyFile string, jwt []byte) ([]byte, error) {
	privKey, err := readPrivateKeyFile(pathPrivateKeyFile)
	if err != nil {
		return nil, err
	}

	publicKey := &privKey.PublicKey

	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.RSA_OAEP, Key: publicKey}, nil)
	if err != nil {
		return nil, err
	}

	object, err := encrypter.Encrypt(jwt)
	if err != nil {
		return nil, err
	}

	serialized, err := object.CompactSerialize()
	if err != nil {
		return nil, err
	}

	serializedEncoded := base64.StdEncoding.EncodeToString([]byte(serialized))
	return []byte(serializedEncoded), nil
}

// return plain text deciphered
func DecipherJwtWithJwe(pathPrivateKeyFile string, jwtCiphered []byte) ([]byte, error) {
	privKey, err := readPrivateKeyFile(pathPrivateKeyFile)
	if err != nil {
		return nil, err
	}

	serializedDecoded, err := base64.StdEncoding.DecodeString(string(jwtCiphered))
	if err != nil {
		return nil, err
	}
	object, err := jose.ParseEncrypted(string(serializedDecoded))
	if err != nil {
		return nil, err
	}

	decrypted, err := object.Decrypt(privKey)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

