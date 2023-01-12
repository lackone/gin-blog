package app

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/pkg/util"
	"time"
)

// jwt载荷
type JwtClaims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.RegisteredClaims
}

// jwt密钥
func GetJwtSecret() []byte {
	return []byte(global.JwtSetting.Secret)
}

// 生成token
func MakeJwt(AppKey, AppSecret string) (string, error) {
	claim := JwtClaims{
		AppKey:    util.EncodeMd5(AppKey),
		AppSecret: util.EncodeMd5(AppSecret),
		RegisteredClaims: jwt.RegisteredClaims{
			//过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(global.JwtSetting.Expire)),
			//签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			//签发人
			Issuer: global.JwtSetting.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(GetJwtSecret())
	return tokenString, err
}

// 解析token
func ParseJwt(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		claims, ok := token.Claims.(*JwtClaims)
		if ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
