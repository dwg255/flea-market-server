package user

import (
	"database/sql"
	"flea-market/common/tools"
	"flea-market/model/userModel"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
var jwtKey = []byte("flea-market")


type Claims struct {
	UserId uint
	NickName string
	Avatar string
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
				c.JSON(200, gin.H{"token": tools.CreateToken(user)})
				return
			}
		} else {
			fmt.Println("find user err: ",err)
		}
	} else {
		c.JSON(200, gin.H{"token": tools.CreateToken(user)})
		return
	}
	c.JSON(500, gin.H{"code": "500","message":"login failed!"})
}
