package utils

import (
	"BiliBili.com/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

var jwtKey = []byte(viper.GetString("server.jwtSecret"))

type Claims struct {
	UserId uint
	Authority int
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) //token过期时间
	claims := &Claims{
		UserId: user.ID,
		Authority:user.Authority,
		StandardClaims: jwt.StandardClaims{
			//发放时间等
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "FanOne",
			Subject:   "token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseUserToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	return token, claims, err
}
