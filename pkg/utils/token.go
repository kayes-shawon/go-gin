package utils

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/oklog/ulid/v2"
	"github.com/pascaldekloe/jwt"
	"time"
)

func GetJTIId() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func GetKeyBytes(key string) []byte {
	data, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil
	}

	return data
}

func GetEdDSAPrivateKey(data []byte) ed25519.PrivateKey {
	block, _ := pem.Decode(data)
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	return key.(ed25519.PrivateKey)
}

func GetEdDSAPublicKey(data []byte) ed25519.PublicKey {
	block, _ := pem.Decode(data)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}

	return key.(ed25519.PublicKey)
}


func Encode(data map[string]interface{}) (token []byte, err error){

	claims := &jwt.Claims{}
	claims.Issuer = "core"
	claims.Audiences = []string{"core"}
	claims.NotBefore = jwt.NewNumericTime(time.Now().Add(-time.Second * 5).Round(time.Second))
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Expires = jwt.NewNumericTime(time.Now().Add(600*time.Second).Round(time.Second))
	claims.ID = GetJTIId()
	claims.Set = data
	privateKey := GetEdDSAPrivateKey(GetKeyBytes("LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1DNENBUUF3QlFZREsyVndCQ0lFSUpVdWRzaCs5c1dGckNFdkJxYmxTYndTbmVXb2VZN2l0QlRRUWI0MHFhTS8KLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLQo="))

	return claims.EdDSASign(privateKey)

}

func RefreshTokenEncode(data map[string]interface{}) ([]byte, error) {
	claims := &jwt.Claims{}
	claims.Issuer = "upay"
	claims.Audiences = []string{"cdfs"}
	claims.NotBefore = jwt.NewNumericTime(time.Now().Add(-time.Second * 5).Round(time.Second))
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Expires = jwt.NewNumericTime(time.Now().Add(3600*time.Second).Round(time.Second))
	claims.ID = GetJTIId()
	claims.Set = data
	privateKey := GetEdDSAPrivateKey(GetKeyBytes("LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1DNENBUUF3QlFZREsyVndCQ0lFSUpVdWRzaCs5c1dGckNFdkJxYmxTYndTbmVXb2VZN2l0QlRRUWI0MHFhTS8KLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLQo="))
	return claims.EdDSASign(privateKey)
}

//func Decode(token string) (bool error) {
//	publicKey := GetEdDSAPublicKey(GetKeyBytes("LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUNvd0JRWURLMlZ3QXlFQUlrSm05OFpJaWdtaEY2RC9LYU9xZXMzSW83S0tHNWExTi9yODJvVGtvWEE9Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="))
//	claims, err := jwt.EdDSACheck([]byte(token), publicKey)
//	if err != nil {
//		return false, err
//	}
//
//	if !claims.Valid(time.Now()) {
//		return false, nil
//	}
//
//	if claims.Issuer != "core" {
//		return false, nil
//	}
//	return true, nil
//}
