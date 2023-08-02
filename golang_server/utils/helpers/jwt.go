package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateAccessToken(username string) (string, error) {
	issueTime := time.Now().Unix()
	expiryTime := issueTime + viper.GetInt64("jwt.accessTokenExpiry")
	claims := customClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issueTime,
			ExpiresAt: expiryTime,
			Issuer:    "estale_ecommerce tutorial",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString("jwt.accessTokenSecret")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt.accessTokenSecret")), nil
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "Invalid JWT access token !")
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return "", errors.New("Couldn't parse token claims !")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return "", errors.New("JWT access token is expired !")
	}
	return claims.Username, nil
}
