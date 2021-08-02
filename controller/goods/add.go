package goods

import (
	"encoding/json"
	"flea-market/common/tools"
	"flea-market/model/goodsModel"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
type AddParams struct {
	Title string `form:"title" json:"title" binding:"required"`
	Price float64 `form:"goods_price" json:"goods_price" binding:"required"`
	Pics []string `form:"pics" json:"pics" binding:"required"`
	UserId int
	Nickname string
	AvatarUrl string
	Address string `form:"address" json:"address" binding:"required"`
	Latitude float64 `form:"latitude" json:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" json:"longitude" binding:"required"`
	Tags []string `form:"goods_tags" json:"goods_tags"`
	Content string `form:"content" json:"content" binding:"required"`
	GoodsNum int `form:"goods_num" json:"goods_num" binding:"required"`
	CatId int `form:"cat_id" json:"cat_id" binding:"required"`
	OnlineSell bool `form:"online_sell" json:"online_sell"`
	ExpressType int `form:"express_type" json:"express_type" binding:"required"`
}

func Add(c *gin.Context) {
	claims := tools.CheckToken(c)
	var addParams AddParams
	if err := c.Bind(&addParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println(addParams)
	pics,_ := json.Marshal(addParams.Pics)
	tags,_ := json.Marshal(addParams.Tags)
	goods := &goodsModel.Goods{
		Title:addParams.Title,
		Price:addParams.Price,
		Pics:string(pics),
		UserId:claims.UserId,
		Nickname:claims.NickName,
		AvatarUrl:claims.Avatar,
		Address:addParams.Address,
		Latitude:addParams.Latitude,
		Longitude:addParams.Longitude,
		Tags:string(tags),
		Content:addParams.Content,
		GoodsNum:addParams.GoodsNum,
		OnlineSell: addParams.OnlineSell,
		ExpressType: addParams.ExpressType,
		CatId:addParams.CatId,
	}

	if _,err := goodsModel.AddGoods(goods); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "发布成功！"})
		return
	}
}