package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateJWT creates an Apple MapKit-compliant JWT string, given a pem key (usually the raw content of a `.p8` file),
// the keyID (10-character key identifier), the teamID (10-character Apple Developer Team ID)
// and an expiration time.
func GenerateJWT(key []byte, keyID string, teamID string, expiry time.Time) (string, error) {
	ecdsaPrivateKey, err := jwt.ParseECPrivateKeyFromPEM(key)
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims := &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiry.Unix(),
		Issuer:    teamID,
	}

	t := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		claims,
	)
	t.Header = map[string]interface{}{
		"alg": jwt.SigningMethodES256.Alg(),
		"kid": keyID,
		"typ": "JWT",
	}
	return t.SignedString(ecdsaPrivateKey)
}
