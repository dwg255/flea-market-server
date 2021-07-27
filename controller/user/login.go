package user

import (
	"database/sql"
	"flea-market/model/userModel"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"time"

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
var jwtkey = []byte("flea-market")
var str string

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	fmt.Println("in login")
	var form LoginParams
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//fmt.Println(form)
	// 根据openid查找 user
	if user, err := userModel.GetUserByOpenId(form.OpenId); err != nil {
		if err == sql.ErrNoRows {
			var gender int
			gender, _ = strconv.Atoi(form.Gender)
			user = &userModel.User{
				Openid:form.OpenId,
				Nickname:form.NickName,
				Gender: gender,
				AvatarUrl:form.AvatarUrl,
				Country:form.Country,
				Province:form.Province,
				City:form.City,
			}
			if user,err = userModel.AddUser(user); err != nil {
				fmt.Println("add user err: ",err)
			} else {
				c.JSON(200, gin.H{"token": createToken(user)})
				return
			}
		} else {
			fmt.Println("find user err: ",err)
		}
	} else {
		fmt.Println(user)
	}
	c.JSON(500, gin.H{"code": "500","message":"login failed!"})
}

func createToken(u *userModel.User) string {
	claims := &Claims{
		UserId: uint(u.UserId),
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
	return tokenString


}
