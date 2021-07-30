package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)
var jwtKey = []byte("flea-market")
type Claims struct {
	UserId int
	jwt.StandardClaims
}
func createToken(userId int) string {
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: expireTime.Unix(), //不设置过期时间
			IssuedAt: time.Now().Unix(),
			Issuer:   "flea market",  // 签名颁发者
			Subject:  "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}
