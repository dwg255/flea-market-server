package goodsModel

import (
	"database/sql"
	"fmt"
	"strconv"

	"flea-market/model"
	"time"
)

type Goods struct {
	GoodsId int `json:"goods_id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Pics string `json:"pics"`
	UserId int `json:"user_id"`
	Nickname string `json:"nickname"`
	AvatarUrl string `json:"avatar_url"`
	Address string `json:"address"`
	Latitude float64 `json:"latitude"`
	Longitude float64`json:"longitude"`
	Tags string`json:"tags"`
	Content string`json:"content"`
	NewMessage interface{}`json:"new_message"`
	GoodsNum int`json:"goods_num"`
	StarNum int`json:"star_num"`
	FavNum int`json:"fav_num"`
	ViewsNum int`json:"views_num"`
	Created int`json:"created"`
	OnlineSell bool	`json:"online_sell"`		// 是否在线交易
	ExpressType int	`json:"express_type"`		// 1 快递 2 自提
	CatId int		`json:"cat_id"`		// 分类ID
	Status int		`json:"status"`		// 状态 0 正常 1 售出 2 下架
}

//查找数量
func GetCount(where string) (int, error) {
	sql := `select count(1) from f_goods ` + where
	//fmt.Println(sql)
	stmt, err :=  model.Db.Prepare(sql)
	if err != nil {
		return 0 ,nil
	}
	defer stmt.Close()
	row := stmt.QueryRow()
	var count int
	err = row.Scan(&count)
	return count ,err
}

// 添加商品
func AddGoods(goods *Goods)(*Goods,error){
	sqlStr := `insert into f_goods (title,price,pics,user_id,nickname,avatar_url,address,latitude,longitude,tags,content,new_message,goods_num,star_num,fav_num,views_num,created,online_sell,express_type,cat_id) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		//fmt.Println("err1",err.Error())
		return nil,err
	} else {
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(goods.Title, goods.Price, goods.Pics, goods.UserId, goods.Nickname, goods.AvatarUrl, goods.Address, goods.Latitude, goods.Longitude, goods.Tags, goods.Content, goods.NewMessage, goods.GoodsNum, goods.StarNum, goods.FavNum, goods.ViewsNum,  time.Now().Unix(),goods.OnlineSell,goods.ExpressType,goods.CatId)
		if err != nil {
			return nil,err
		}

		lastInsertId,_ := res.LastInsertId()
		goods.UserId = int(lastInsertId)
		return goods,err
	}
}

// 编辑商品
func UpdateStarNum(goodsId int,num int)(int64,error){
	sqlStr := "update f_goods set star_num = star_num + ? where goods_id = ?"
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		fmt.Println("err1",err.Error())
		return 0,err
	} else {
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(num,goodsId)
		if err != nil {
			return 0,err
		}
		rowsAffected,_ := res.RowsAffected()
		if rowsAffected == 0 {
			err = nil	//数据没有修改
		}
		return rowsAffected,err
	}
}

// 编辑商品
func UpdateFavNum(goodsId int,num int)(int64,error){
	sqlStr := "update f_goods set fav_num = fav_num + ? where goods_id = ?"
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		fmt.Println("err1",err.Error())
		return 0,err
	} else {
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(num,goodsId)
		if err != nil {
			return 0,err
		}
		rowsAffected,_ := res.RowsAffected()
		if rowsAffected == 0 {
			err = nil	//数据没有修改
		}
		return rowsAffected,err
	}
}

// 编辑商品
func UpdateStatus(goodsId,userId,status int)(int64,error){
	sqlStr := "update f_goods set status = ? where goods_id = ? and user_id = ?"
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		fmt.Println("err1",err.Error())
		return 0,err
	} else {
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(status,goodsId,userId)
		if err != nil {
			return 0,err
		}
		rowsAffected,_ := res.RowsAffected()
		return rowsAffected,err
	}
}

// 编辑商品
func UpdateGoods(goods *Goods)(*Goods,error){
	sqlStr := "update f_goods set `title` = ?, `price` = ? ,pics = ?,user_id = ?,`nickname` = ?,avatar_url = ?,`address` = ?,`latitude` = ?,`longitude` = ?,`tags` = ?,`content` = ?,new_message = ?,goods_num = ?,star_num = ?,fav_num = ?,views_num = ?,online_sell = ?,express_type = ?, cat_id = ? where goods_id = ?"
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		fmt.Println("err1",err.Error())
		return nil,err
	} else {
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(goods.Title, goods.Price, goods.Pics, goods.UserId, goods.Nickname, goods.AvatarUrl, goods.Address, goods.Latitude, goods.Longitude, goods.Tags, goods.Content, goods.NewMessage, goods.GoodsNum, goods.StarNum, goods.FavNum, goods.ViewsNum, goods.OnlineSell,goods.ExpressType,goods.CatId,goods.GoodsId)
		if err != nil {
			return nil,err
		}

		rowsAffected,_ := res.RowsAffected()
		if rowsAffected == 0 {
			err = nil	//数据没有修改
			//fmt.Println("数据没有修改ss")
		}
		//fmt.Println("rowsAffected = " ,rowsAffected)
		return goods,err
	}
}

// 改变状态


// 根据主键查找
func GetGoodsById(goodsId int) (goods *Goods, err error) {
	sql := `select * from f_goods where goods_id = ?`
	stmt, err :=  model.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(goodsId)
	goods,err = ScanRow(row)
	return
}

// 条件查找
func GetGoods(where string,offset int,limit int) (goodsList []*Goods, err error) {
	sqlStr := `select * from f_goods ` + where + ` limit ?,?`
	var stmt *sql.Stmt
	stmt, err =  model.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(sqlStr)
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	rows,err := stmt.Query(offset,limit)
	if err != nil {
		return
	}
	var goods *Goods
	goodsList = make([]*Goods,0)
	//fmt.Println(sqlStr,offset ,limit )
	for rows.Next() {
		goods,err = ScanRows(rows)
		if err != nil {
			return
		}
		goodsList = append(goodsList, goods)
	}
	return
}

// 根据用户id查找
func GetGoodsByUserId(userId int,offset int,limit int) (goodsList []*Goods, err error) {
	return GetGoods("user_id=" + strconv.Itoa(userId),offset ,limit )
}

func ScanRow (row *sql.Row) (u *Goods,err error) {
	var (
		goodsId int
		title string
		price float64
		pics string
		userId int
		nickname string
		avatarUrl string
		address string
		latitude float64
		longitude float64
		tags string
		content string
		newMessage string
		goodsNum int
		starNum int
		favNum int
		viewsNum int
		created int
		onlineSell bool
		expressType int
		catId int
		status int
	)
	if err = row.Scan(&goodsId,&title,&price,&pics,&userId,&nickname,&avatarUrl,&address,&latitude,&longitude,&tags,&content,&newMessage,&goodsNum,&starNum,&favNum,&viewsNum,&created,&onlineSell,&expressType,&catId,&status);err != nil {
		return
	}
	u = &Goods{
		goodsId,
		title,
		price,
		pics,
		userId,
		nickname,
		avatarUrl,
		address,
		latitude,
		longitude,
		tags,
		content,
		newMessage,
		goodsNum,
		starNum,
		favNum,
		viewsNum,
		created,
		onlineSell,
		expressType,
		catId,
		status,
	}
	return
}

func ScanRows (row *sql.Rows) (u *Goods,err error) {
	var (
		goodsId int
		title string
		price float64
		pics string
		userId int
		nickname string
		avatarUrl string
		address string
		latitude float64
		longitude float64
		tags string
		content string
		newMessage string
		goodsNum int
		starNum int
		favNum int
		viewsNum int
		created int
		onlineSell bool
		expressType int
		catId int
		status int
	)
	if err = row.Scan(&goodsId,&title,&price,&pics,&userId,&nickname,&avatarUrl,&address,&latitude,&longitude,&tags,&content,&newMessage,&goodsNum,&starNum,&favNum,&viewsNum,&created,&onlineSell,&expressType,&catId,&status);err != nil {
		return
	}
	u = &Goods{
		goodsId,
		title,
		price,
		pics,
		userId,
		nickname,
		avatarUrl,
		address,
		latitude,
		longitude,
		tags,
		content,
		newMessage,
		goodsNum,
		starNum,
		favNum,
		viewsNum,
		created,
		onlineSell,
		expressType,
		catId,
		status,
	}
	return
}