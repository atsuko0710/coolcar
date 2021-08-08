package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenGenerator struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFuc     func() time.Time
}

func NewJWTTokenGenerator(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGenerator {
	return &JWTTokenGenerator{
		issuer:     issuer,
		nowFuc:     time.Now,
		privateKey: privateKey,
	}
}

func (j *JWTTokenGenerator) GenerateToken(accountID string, expire time.Duration) (string, error) {
	nowSec := j.nowFuc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Subject:   accountID,
		Issuer:    j.issuer,
		IssuedAt:  nowSec,
		ExpiresAt: nowSec + int64(expire.Seconds()),
	})

	return tkn.SignedString(j.privateKey)
}
