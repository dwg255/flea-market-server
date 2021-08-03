package star

import (
	"flea-market/common/tools"
	"flea-market/model/goodsModel"
	"flea-market/model/starModel"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddParams struct {
	UserId int `form:"user_id" json:"user_id" binding:"required"`
	GoodsId int `form:"goods_id" json:"goods_id" binding:"required"`
}

func Add(c *gin.Context) {
	claims := tools.CheckToken(c)
	var addParams AddParams
	if err := c.Bind(&addParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		if _,err := starModel.GetStar(claims.UserId,addParams.GoodsId); err != nil {
			if affectedRows,err := starModel.AddStar(&starModel.Star{UserId:claims.UserId,GoodsId: addParams.GoodsId});err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			} else {
				if affectedRows > 0 {
					c.JSON(http.StatusOK, gin.H{"msg": "点赞成功！"})
					goodsModel.UpdateStarNum(addParams.GoodsId,1)
					return
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{"msg": "点赞成功！"})
		return
	}
}