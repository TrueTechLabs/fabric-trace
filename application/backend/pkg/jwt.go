package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type Myclaims struct {
	UserID   string `json:"userID"`
	UserType string `json:"userType"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 48

var MySecret = []byte(viper.GetString("jwt.secret"))

func GenToken(userID string, userType string) (string, error) {
	c := Myclaims{
		UserID:   userID,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    viper.GetString("jwt.issuer"),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

// 解析token
func ParseToken(tokenString string) (*Myclaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Myclaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Myclaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
