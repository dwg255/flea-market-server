package goods

import (
	"flea-market/model/goodsModel"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 定义接收数据的结构体
type ListParams struct {
	CatId int `form:"cat_id" json:"cat_id" uri:"cat_id" xml:"cat_id"`
	PageSize int `form:"page_size" json:"page_size" uri:"page_size" xml:"page_size" binding:"required"`
	PageNum int `form:"page_num" json:"page_num" uri:"page_num" xml:"page_num" binding:"required"`
}

func List(c *gin.Context) {
	var listParams ListParams
	if err := c.Bind(&listParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println(listParams)
	where := "where status = 0 "
	if listParams.CatId != 0 {
		where += " and cat_id = " + strconv.Itoa(listParams.CatId)
	}

	if count,err := goodsModel.GetCount(where); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":"err"})
		return
	} else {
		if count == 0 {
			c.JSON(http.StatusOK,gin.H{"total":0})
		} else {
			//fmt.Println(count)
			if listParams.PageSize * (listParams.PageNum - 1) >= count {
				c.JSON(http.StatusOK,gin.H{"msg":"无更多内容"})
				return
			}
			if list,err := goodsModel.GetGoods(where , listParams.PageSize * (listParams.PageNum - 1), listParams.PageSize);err != nil {
				c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
				return
			} else {
				c.JSON(http.StatusOK,gin.H{"list":list,"total":count,"page_size":listParams.PageSize})
			}
		}
	}

}