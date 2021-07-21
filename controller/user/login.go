package user

import (
	"flea-market/model"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type LoginParams struct {
	OpenId    string `form:"openid" json:"openid" binding:"required"`
	NickName  string `form:"nickName" json:"nickName" binding:"required"`
	Gender    string `form:"gender" json:"gender"`
	AvatarUrl string `form:"avatarUrl" json:"avatarUrl"`
	Country   string `form:"country" json:"country"`
	Province  string `form:"province" json:"province"`
	City      string `form:"city" json:"city"`
}

//自定义一个字符串
var jwtkey = []byte("www.topgoer.com")
var str string

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var form LoginParams
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(form)
	// 根据openid查找 user
	if user, err := model.GetUserByOpenId(form.OpenId); err != nil {
		log.Fatal(err)
	} else {
		log.Fatal(user)
	}
	claims := &Claims{
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: expireTime.Unix(), //不设置过期时间
			IssuedAt: time.Now().Unix(),
			Issuer:   "127.0.0.1",  // 签名颁发者
			Subject:  "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	str = tokenString
	c.JSON(200, gin.H{"token": tokenString})
}
