package goods

import (
	"flea-market/common/tools"
	"flea-market/model/goodsModel"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateStatus(c *gin.Context) {
	claims := tools.CheckToken(c)
	goodsIdStr := c.Query("goods_id")
	statusStr := c.Query("status")
	if status,err := strconv.Atoi(statusStr);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
		return
	} else {
		if goodsId,err := strconv.Atoi(goodsIdStr);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
			return
		} else {
			if affectedRows, err := goodsModel.UpdateStatus(status,goodsId,claims.UserId); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			} else {
				if affectedRows > 0 {
					c.JSON(http.StatusOK, gin.H{"msg": "下架成功！"})
					return
				}
				c.JSON(http.StatusBadRequest, gin.H{"error": "失败"})
			}
		}
	}

}