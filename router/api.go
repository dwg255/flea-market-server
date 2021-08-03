package router

import (
	"flea-market/controller/dialog"
	"flea-market/controller/goods"
	"flea-market/controller/star"
	"flea-market/controller/user"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		// 用户相关
		api.GET("/user/login", user.Login)

		// 商品相关
		api.POST("/goods/add",goods.Add)
		api.POST("/goods/update",goods.Update)
		api.POST("/goods/list",goods.List)
		api.POST("/goods/detail",goods.Detail)
		api.POST("/goods/status",goods.UpdateStatus)

		// 留言相关
		api.POST("/dialog/add",dialog.Add)
		api.POST("/dialog/list",dialog.List)
		api.POST("/dialog/response",dialog.Response)

		// 点赞相关
		api.POST("/star/add",star.Add)
		api.POST("/star/remove",star.Remove)
	}
}
