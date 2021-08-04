package goods

import (
	"database/sql"
	"flea-market/common/tools"
	"flea-market/model/dialogModel"
	"flea-market/model/goodsModel"
	"flea-market/model/starModel"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Detail(c *gin.Context) {
	claims := tools.CheckToken(c)
	goodsIdStr := c.Query("goods_id")
	//fmt.Println("goods_id = ",goodsIdStr)
	if goodsId,err := strconv.Atoi(goodsIdStr);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":err})
	} else {
		if goods,err := goodsModel.GetGoodsById(goodsId);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
		} else {
			if dialogList,err := dialogModel.GetDialogs(" where goods_id = " + goodsIdStr); err != nil {
				c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
			} else {
				var isStar bool
				if _,err := starModel.GetStar(claims.UserId,goodsId); err != nil {
					if err == sql.ErrNoRows {

					}
					//return
				} else {
					isStar = true
				}

				total,_ := goodsModel.GetCount(" where user_id = " + strconv.Itoa(goods.GoodsId))

				c.JSON(http.StatusOK,gin.H {
					"goodsInfo":goods,
					"dialogList":dialogList,
					"star":isStar,
					"goods_num":total,
				})
				//fmt.Println("[before] ",goods)
				goods.ViewsNum ++
				 goodsModel.UpdateGoods(goods)
				//fmt.Println("[after] ",g)
				//fmt.Println(e)
			}
		}
	}
}