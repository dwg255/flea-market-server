package goods

import (
	"encoding/json"
	"flea-market/model/goodsModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
type AddParams struct {
	Title string `form:"title" binding:"required"`
	Price float64 `form:"goods_price" binding:"required"`
	Pics []string `form:"pics" binding:"required"`
	UserId int
	Nickname string
	AvatarUrl string
	Address string `form:"address" binding:"required"`
	Latitude float64 `form:"latitude" binding:"required"`
	Longitude float64 `form:"latitude" binding:"required"`
	Tags []string
	Content string `form:"content" binding:"required"`
	GoodsNum int `form:"goods_num" binding:"required"`
	CatId int `form:"cat_id" binding:"required"`
}
func Add(c *gin.Context) {
	var addParams AddParams
	if err := c.Bind(&addParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pics,_ := json.Marshal(addParams.Pics)
	tags,_ := json.Marshal(addParams.Tags)
	goods := &goodsModel.Goods{
		Title:addParams.Title,
		Price:addParams.Price,
		Pics:string(pics),
		UserId:0,
		Nickname:"",
		AvatarUrl:addParams.AvatarUrl,
		Address:addParams.Address,
		Latitude:addParams.Latitude,
		Longitude:addParams.Longitude,
		Tags:string(tags),
		Content:addParams.Content,
		GoodsNum:addParams.GoodsNum,
	}
	if _,err := goodsModel.AddGoods(goods); err != nil {
		fmt.Println("add user err: ",err)
	} else {
		return
	}
}