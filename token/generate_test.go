package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

var key = []byte("-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgKlW/I1ah7UnJl7JK\nfwRFR90NYiwm16GVWaDhoC2S9iyhRANCAAShnP/JD1CrVYsIZbonaT94kLEc2NaL\n0ACA/1RbM9XP0LeVOJxYddbAheJ90Zv+tNCgHR7Qa8X6CARr3aG1h+QB\n-----END PRIVATE KEY-----")
var invalid = []byte("-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgKlW/I1ah7UnJl7JK\nfwRFR90NYiwm16GVWaDhoC2S9iyhRANCAAShnP/JD1CrVYsIZbonaT94kLEc2NaL\n0ACA/1R\n-----END PRIVATE KEY-----")

func TestGenerateJWT(t *testing.T) {
	now := time.Now()
	expiry := now.Add(time.Hour)
	token, err := GenerateJWT(key, "1234567890", "ABCD123456", expiry)
	assert.NoError(t, err)

	ecdsaPrivateKey, _ := jwt.ParseECPrivateKeyFromPEM(key)
	tok, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return ecdsaPrivateKey.Public(), nil
	})
	assert.NoError(t, err)
	assert.Equal(t, "1234567890", tok.Header["kid"])
	assert.Equal(t, "JWT", tok.Header["typ"])
	assert.Equal(t, "ES256", tok.Header["alg"])

	claims := tok.Claims.(*jwt.StandardClaims)
	assert.Equal(t, "ABCD123456", claims.Issuer)
	assert.Equal(t, now.Unix(), claims.IssuedAt)
	assert.Equal(t, expiry.Unix(), claims.ExpiresAt)

}

func TestGenerateJWT_Invalid(t *testing.T) {
	now := time.Now()
	expiry := now.Add(time.Hour)
	_, err := GenerateJWT(invalid, "1234567890", "ABCD123456", expiry)
	assert.Error(t, err)
}
