package dialog

import (
	"flea-market/common/tools"
	"flea-market/model/dialogModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
type Res struct {
	err error
	data gin.H
}

func Response(c *gin.Context) {
	claims := tools.CheckToken(c)
	var res = Res{
		nil,gin.H{},
	}
	defer func() {
		if res.err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"msg":res.err.Error()})
		} else {
			c.JSON(http.StatusOK,res.data)
		}
	}()
	idStr := c.Query("id")
	userIdStr := c.Query("user_id")
	message := c.Query("message")
	//fmt.Println("goods_id = ",goodsIdStr)
	if userId,err := strconv.Atoi(userIdStr);err != nil {
		res.err = err
		//c.JSON(http.StatusBadRequest,gin.H{"msg":err})
		return
	} else {
		if claims.UserId != userId {
			//c.JSON(http.StatusBadRequest,gin.H{"msg":err})
			res.err = fmt.Errorf("只有卖家可以回复")
			return
		}
	}
	if id,err := strconv.Atoi(idStr);err != nil {
		//c.JSON(http.StatusBadRequest,gin.H{"msg":err})
		res.err = err
		return
	} else {
		if affectedRows,err := dialogModel.UpdateDialog(message," where id = "+ strconv.Itoa(id));err != nil {
			res.err = err
			return
		} else {
			if affectedRows != 0 {
				res.data = gin.H{"msg":"回复成功！"}
				return
			} else {
				res.err = fmt.Errorf("回复失败")
				return
			}
		}
	}
}