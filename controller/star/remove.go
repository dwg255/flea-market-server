package star

import (
	"database/sql"
	"flea-market/common/tools"
	"flea-market/model/goodsModel"
	"flea-market/model/starModel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Remove(c *gin.Context) {
	claims := tools.CheckToken(c)
	var addParams AddParams
	if err := c.Bind(&addParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		if star,err := starModel.GetStar(claims.UserId,addParams.GoodsId); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"msg": "取消点赞成功！"})
			}
			return
		} else {
			if affectedRows,err := starModel.RemoveStar(star);err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			} else {
				if affectedRows > 0 {
					c.JSON(http.StatusOK, gin.H{"msg": "取消点赞成功！"})
					goodsModel.UpdateStarNum(addParams.GoodsId,-1)
					return
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{"msg": "点赞成功！"})
		return
	}
}