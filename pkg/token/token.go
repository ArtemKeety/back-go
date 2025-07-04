package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	AccessTime  = 3600 //seconds
	RefreshTime = 24   // hours
	Issuer      = "Artem Medvedev"
)

var secret = []byte("Artem")

type AccessToken struct {
	jwt.RegisteredClaims
	Guid string `json:"guid"`
}

type RefreshToken struct {
	jwt.RegisteredClaims
	Guid string
}

func NewAccessToken(guid string) (string, error) {
	claims := AccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTime * time.Second)),
			Issuer:    Issuer,
		},
		Guid: guid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewRefreshToken(guid string) (string, error) {
	claims := RefreshToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: Issuer,
		},
		Guid: guid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*AccessToken, error) {}
