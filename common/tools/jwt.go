package tools

import (
	"flea-market/model/userModel"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
var jwtKey = []byte("flea-market")
type Claims struct {
	UserId int
	NickName string
	Avatar string
	jwt.StandardClaims
}

func CreateToken(u *userModel.User) string {
	claims := &Claims{
		UserId: u.UserId,
		NickName: u.Nickname,
		Avatar: u.AvatarUrl,
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

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, Claims, err
}

func CheckToken(c *gin.Context) *Claims {
	tokenString := c.GetHeader("token")
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return nil
	}
	return claims
}