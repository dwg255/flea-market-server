package goods

import (
	"flea-market/model/dialogModel"
	"flea-market/model/goodsModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 定义接收数据的结构体
type ListParams struct {
	CatId int `form:"cat_id" json:"cat_id" uri:"cat_id" xml:"cat_id"`
	UserId int `form:"shop_id" json:"shop_id" uri:"shop_id" xml:"shop_id"`
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
	//where := "where status = 0 "
	whereMap := make(map[string]string)
	if listParams.CatId != 0 {
		whereMap["cat_id"] = strconv.Itoa(listParams.CatId)
	}
	if listParams.UserId != 0 {
		whereMap["user_id"] = strconv.Itoa(listParams.UserId)
	}
	where := ""
	args := make([]interface{},0)
	if len(whereMap) != 0 {
		where = "where "
		whereArr := make([]string,0)
		for k,v := range whereMap {
			if k == "keywords" {
				whereArr = append(whereArr, fmt.Sprintf(" title like ? or content like ? " ))
				args = append(args, "%"+ v + "%","%"+ v + "%")
			} else {
				whereArr = append(whereArr, k + " = ? ")
				args = append(args, v )
			}
		}
		for i := range whereArr {
			if i == len(whereArr)-1 {
				where += whereArr[i]
			} else {
				where += whereArr[i] + " and "
			}
		}
	}

	//where := ""
	if listParams.CatId != 0 {
		where += " where cat_id = " + strconv.Itoa(listParams.CatId)
	}

	if count,err := goodsModel.GetCount(where,args...); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
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
			where += " limit ?,? "
			args = append(args, listParams.PageSize * (listParams.PageNum - 1), listParams.PageSize)
			if list,err := goodsModel.GetGoods(where ,args... );err != nil {
				c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
				return
			} else {
				if len(list) >0 {
					type res struct {
						goodsModel.Goods
						dialogModel.Dialog
					}
					where := " where id in (select max(id) from f_dialog group by goods_id having goods_id in ("
					for i:= 0; i< len(list);i++ {
						where += strconv.Itoa(list[i].GoodsId) + ","
					}
					where = where[:len(where)-1] + ")) "
					//fmt.Println(where)
					if dialogList,err := dialogModel.GetDialogs(where);err != nil {
						c.JSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
						return
					} else {
						for i := range list {
							for j := range dialogList {
								if list[i].GoodsId == dialogList[j].GoodsId {
									list[i].NewMessage = dialogList[j]
									break
								}
							}
						}
					}

					c.JSON(http.StatusOK,gin.H{"list":list,"total":count,"page_size":listParams.PageSize})
					return
				}
			}
		}
	}

}