package app

import (
	"my-blog-sevice/global"
	"my-blog-sevice/pkg/util"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expires"`
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(username, password string) (JWTOutput, error) {
	nowTime := time.Now()
	expiretime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		UserName: util.EncodeMD5(username),
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiretime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	jwtOutput := JWTOutput{
		Token:  token,
		Expire: expiretime,
	}
	return jwtOutput, err
}

func ParseToken(token string) (*Claims, error) {
	claims := new(Claims)
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}

	}
	return nil, err
}
