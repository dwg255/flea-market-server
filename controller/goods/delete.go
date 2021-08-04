package goods


import (
	"flea-market/common/tools"
	"flea-market/model/goodsModel"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func Delete(c *gin.Context) {
	claims := tools.CheckToken(c)
	goodsIdStr := c.Query("goods_id")
	if goodsId,err := strconv.Atoi(goodsIdStr);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
		return
	} else {
		if _,err := goodsModel.UpdateStatus(goodsId,claims.UserId,2);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
			return
		} else {
			c.JSON(http.StatusOK,gin.H{"msg":"删除成功！"})
		}
	}

}