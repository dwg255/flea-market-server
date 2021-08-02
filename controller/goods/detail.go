package goods

import (
	"flea-market/model/dialogModel"
	"flea-market/model/goodsModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Detail(c *gin.Context) {
	goodsIdStr := c.Query("goods_id")
	fmt.Println("goods_id = ",goodsIdStr)
	if goodsId,err := strconv.Atoi(goodsIdStr);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":err})
	} else {
		if goods,err := goodsModel.GetGoodsById(goodsId);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
		} else {
			if dialogList,err := dialogModel.GetDialogs(" where goods_id = " + goodsIdStr); err != nil {
				c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
			} else {
				c.JSON(http.StatusOK,gin.H {
					"goodsInfo":goods,
					"dialogList":dialogList,
				})
			}
		}
	}
}