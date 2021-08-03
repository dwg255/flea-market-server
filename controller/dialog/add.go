package dialog

import (
	"flea-market/common/tools"
	"flea-market/model/dialogModel"
	"flea-market/model/goodsModel"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 定义接收数据的结构体
type AddParams struct {
	GoodsId int `form:"goods_id" json:"goods_id" binding:"required"`
	Message string `form:"message" json:"message" binding:"required"`
}

func Add(c *gin.Context) {
	claims := tools.CheckToken(c)
	var addParams AddParams
	if err := c.Bind(&addParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if goods,err := goodsModel.GetGoodsById(addParams.GoodsId);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
	} else {
		dialog := &dialogModel.Dialog{
			GoodsId:addParams.GoodsId,
			UserId:goods.UserId,
			Avatar:goods.AvatarUrl,
			Nickname:goods.Nickname,
			CustomerUserId:claims.UserId,
			CustomerAvatar:claims.Avatar,
			CustomerNickname:claims.NickName,
			Question:addParams.Message,
			//Status:0,
			Created: int(time.Now().Unix()),
		}
		if _,err := dialogModel.AddDialog(dialog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "发布成功！"})
			goodsModel.UpdateFavNum(addParams.GoodsId,1)
			return
		}
	}


}