package starModel

import (
	"database/sql"
	"flea-market/model"
)

type Star struct {
	Id int `json:"id"`
	GoodsId int `json:"goods_id"`
	UserId int `json:"user_id"`
}

// 点赞
func AddStar(star *Star)(int64,error){
	sqlStr := `insert into f_star (user_id,goods_id) values(?,?)`
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		//fmt.Println("err1",err.Error())
		return 0,err
	} else {
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(star.UserId,star.GoodsId)
		if err != nil {
			return 0,err
		}
		lastInsertId,_ := res.LastInsertId()
		return lastInsertId,err
	}
}

// 取消点赞
func RemoveStar(star *Star)(int64,error){
	sqlStr := `delete FROM  f_star where user_id=? & goods_id=?`
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		//fmt.Println("err1",err.Error())
		return 0,err
	} else {
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(star.UserId,star.GoodsId)
		if err != nil {
			return 0,err
		}
		RowsAffected,_ := res.RowsAffected()
		return RowsAffected,err
	}
}


// 根据主键查找
func GetStar(userId int,goodsId int) (star *Star, err error) {
	sql := `select * from f_star where user_id = ? and goods_id = ?`
	stmt, err :=  model.Db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId,goodsId)
	star,err = ScanRow(row)
	return

}

func ScanRows (row *sql.Rows) (u *Star,err error) {
	var Id int
	var UserId    int
	var GoodsId int
	if err = row.Scan(&Id, &UserId, &GoodsId);err != nil {
		return
	}
	u = &Star{
		Id,
		UserId,
		GoodsId,
	}
	return
}

func ScanRow (row *sql.Row) (u *Star,err error) {
	var Id int
	var UserId    int
	var GoodsId int
	if err = row.Scan(&Id, &UserId, &GoodsId);err != nil {
		return
	}
	u = &Star{
		Id,
		UserId,
		GoodsId,
	}
	return
}