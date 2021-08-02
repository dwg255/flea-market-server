package router

import (
	"flea-market/controller/goods"
	"flea-market/controller/user"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		//用户相关
		api.GET("/user/login", user.Login)
		//商品相关
		api.POST("/goods/add",goods.Add)
		api.POST("/goods/update",goods.Update)
		api.POST("/goods/list",goods.List)
		api.POST("/goods/detail",goods.Detail)
	}
}
