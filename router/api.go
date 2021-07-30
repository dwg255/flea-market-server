package router

import (
	"flea-market/controller/goods"
	"flea-market/controller/user"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/user/login", user.Login)
		api.POST("/goods/add",goods.Add)
	}
}
