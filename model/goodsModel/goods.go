package goodsModel

import (
	"database/sql"
	"strconv"

	"flea-market/model"
	"time"
)

type Goods struct {
	GoodsId int
	Title string
	Price float64
	Pics string
	UserId int
	Nickname string
	AvatarUrl string
	Address string
	Latitude float64
	Longitude float64
	Tags string
	Content string
	NewMessage string
	GoodsNum int
	StarNum int
	FavNum int
	ViewsNum int
	Created int
}

//查找数量
func GetCount(where string) (int, error) {
	sql := `select count(1) from f_goods ` + where
	stmt, err :=  model.Db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0 ,nil
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var count int
	err = row.Scan(&count)
	return count ,err
}

// 添加用户
func AddGoods(goods *Goods)(*Goods,error){
	sqlStr := `insert into f_goods (title,price,pics,user_id,nickname,avatar_url,address,latitude,longitude,tags,content,new_message,goods_num,star_num,fav_num,views_num,created) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		return nil,err
	} else {
		defer stmt.Close()
		if err != nil {
			return nil,err
		}
		res, err := stmt.Exec(goods.Title, goods.Price, goods.Pics, goods.UserId, goods.Nickname, goods.AvatarUrl, goods.Address, goods.Latitude, goods.Longitude, goods.Tags, goods.Content, goods.NewMessage, goods.GoodsNum, goods.StarNum, goods.FavNum, goods.ViewsNum, goods.Created, time.Now().Unix())
		lastInsertId,_ := res.LastInsertId()
		goods.UserId = int(lastInsertId)
		return goods,err
	}
}

// 根据主键查找
func GetGoodsById(goodsId int) (goods *Goods, err error) {
	sql := `select * from f_goods where goods_id = ?`
	stmt, err :=  model.Db.Prepare(sql)
	defer stmt.Close()
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
	sqlStr := `select * from f_goods ` + where + ` limit(?,?)`
	var stmt *sql.Stmt
	stmt, err =  model.Db.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		return
	}
	defer stmt.Close()
	rows,err := stmt.Query(offset,limit)
	if err != nil {
		return
	}
	var goods *Goods
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
	)
	if err = row.Scan(&goodsId,&title,&price,&pics,&userId,&nickname,&avatarUrl,&address,&latitude,&longitude,&tags,&content,&newMessage,&goodsNum,&starNum,&favNum,&viewsNum,&created,);err != nil {
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
	)
	if err = row.Scan(&goodsId,&title,&price,&pics,&userId,&nickname,&avatarUrl,&address,&latitude,&longitude,&tags,&content,&newMessage,&goodsNum,&starNum,&favNum,&viewsNum,&created,);err != nil {
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
	}
	return
}