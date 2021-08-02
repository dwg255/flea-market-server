package goods

import (
"encoding/json"
"flea-market/common/tools"
"flea-market/model/goodsModel"
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
	//fmt.Println(addParams)
	pics,_ := json.Marshal(updateParams.Pics)
	tags,_ := json.Marshal(updateParams.Tags)
	goods := &goodsModel.Goods{
		GoodsId: updateParams.GoodsId,
		Title:updateParams.Title,
		Price:updateParams.Price,
		Pics:string(pics),
		UserId:claims.UserId,
		Nickname:claims.NickName,
		AvatarUrl:claims.Avatar,
		Address:updateParams.Address,
		Latitude:updateParams.Latitude,
		Longitude:updateParams.Longitude,
		Tags:string(tags),
		Content:updateParams.Content,
		GoodsNum:updateParams.GoodsNum,
		OnlineSell: updateParams.OnlineSell,
		ExpressType: updateParams.ExpressType,
		CatId:updateParams.CatId,
	}

	if _,err := goodsModel.UpdateGoods(goods); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "发布成功！"})
		return
	}
}