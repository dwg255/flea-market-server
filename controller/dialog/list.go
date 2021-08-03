package dialog

import (
	"flea-market/model/dialogModel"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func List(c *gin.Context) {
	goodsIdStr := c.Query("goods_id")
	//fmt.Println("goods_id = ",goodsIdStr)
	if goodsId,err := strconv.Atoi(goodsIdStr);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":err})
	} else {
		if dialogList,err := dialogModel.GetDialogs(" where goods_id = " + strconv.Itoa(goodsId));err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
		} else {
			c.JSON(http.StatusOK,gin.H {
				"dialogList":dialogList,
			})
		}
	}

}