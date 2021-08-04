package goods

import (
	"encoding/json"
	"flea-market/common/tools"
	"flea-market/model/goodsModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
type UpdateParams struct {
	GoodsId int `form:"goods_id" json:"goods_id" binding:"required"`
	AddParams
}

func Update(c *gin.Context) {
	claims := tools.CheckToken(c)

	var updateParams UpdateParams
	if err := c.Bind(&updateParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pics,_ := json.Marshal(updateParams.Pics)
	tags,_ := json.Marshal(updateParams.Tags)
	whereMap := map[string]interface{}{
		"title":updateParams.Title,
		"price":updateParams.Price,
		"pics":pics,
		"tags":tags,
		"content":updateParams.Content,
		"goods_num":updateParams.GoodsNum,
		"online_sell":updateParams.OnlineSell,
		"express_type":updateParams.ExpressType,
		"cat_id":updateParams.CatId,
	}
	sql := "update f_goods set "
	args := make([]interface{},0)
	for k, v := range whereMap {
		sql += fmt.Sprintf(" %s = ?,",k)
		args = append(args, v)
	}
	sql = sql[:len(sql) - 1] + " where user_id = ? and goods_id = ?"
	fmt.Println(sql)
	args = append(args,claims.UserId,updateParams.GoodsId)
	if _,err := goodsModel.Update(sql,args...);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"msg":"修改成功！"})
}