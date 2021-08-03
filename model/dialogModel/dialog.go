package dialogModel

import (
	"database/sql"
	"flea-market/model"
	"fmt"
	"time"
)

type Dialog struct {
	Id int `json:"id"`
	GoodsId int `json:"goods_id"`
	UserId int `json:"user_id"`
	Avatar string `json:"avatar"`
	Nickname string `json:"nickname"`
	CustomerUserId int `json:"customer_user_id"`
	CustomerAvatar string `json:"customer_avatar"`
	CustomerNickname string `json:"customer_nickname"`
	Question string `json:"question"`
	Answer string `json:"answer"`
	Status int `json:"status"`
	Created int `json:"created"`
}

// 添加对话
func AddDialog(dialog *Dialog)(*Dialog,error){
	sqlStr := `insert into f_dialog (goods_id,user_id,avatar,nickname,customer_user_id,customer_avatar,customer_nickname,question,answer,status,created) values(?,?,?,?,?,?,?,?,?,?,?)`
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		//fmt.Println("err1",err.Error())
		return nil,err
	} else {
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(dialog.GoodsId,dialog.UserId,dialog.Avatar,dialog.Nickname,dialog.CustomerUserId,dialog.CustomerAvatar,dialog.CustomerNickname,dialog.Question,dialog.Answer,dialog.Status,time.Now().Unix())
		if err != nil {
			return nil,err
		}

		lastInsertId,_ := res.LastInsertId()
		dialog.Id = int(lastInsertId)
		return dialog,err
	}
}

// 编辑商品
func UpdateDialog(answer string,where string)(int64,error){
	sqlStr := `update f_dialog set answer = ? ` + where
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		fmt.Println("err1",err.Error())
		return 0,err
	} else {
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(answer)
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
// 查询商品全部对话

// 条件查找
func GetDialogs(where string) (dialogList []*Dialog, err error) {
	sqlStr := `select * from f_dialog ` + where
	var stmt *sql.Stmt
	stmt, err =  model.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	rows,err := stmt.Query()
	if err != nil {
		return
	}
	var dialog *Dialog
	dialogList = make([]*Dialog,0)
	//fmt.Println(sqlStr,offset ,limit )
	for rows.Next() {
		dialog,err = ScanRows(rows)
		if err != nil {
			return
		}
		dialogList = append(dialogList, dialog)
	}
	return
}
// 条件查找
func GetOne(where string,order string) (dialog *Dialog, err error) {
	sqlStr := `select * from f_dialog ` + where + order + " limit 1 "
	var stmt *sql.Stmt
	stmt, err =  model.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow()

	dialog,err = ScanRow(row)
	return
}

func ScanRow (row *sql.Row) (dialog *Dialog,err error) {
	var (
		id int
		goodsId int
		userId int
		avatar string
		nickname string
		customerUserId int
		customerAvatar string
		customerNickname string
		question string
		answer string
		status int
		created int
	)
	if err = row.Scan(&id,&goodsId,&userId,&avatar,&nickname,&customerUserId,&customerAvatar,&customerNickname,&question,&answer,&status,&created);err != nil {
		return
	}
	dialog = &Dialog{
		id,
		goodsId,
		userId,
		avatar,
		nickname,
		customerUserId,
		customerAvatar,
		customerNickname,
		question,
		answer,
		status,
		created,
	}
	return
}

func ScanRows (row *sql.Rows) (dialog *Dialog,err error) {
	var (
		id int
		goodsId int
		userId int
		avatar string
		nickname string
		customerUserId int
		customerAvatar string
		customerNickname string
		question string
		answer string
		status int
		created int
	)
	if err = row.Scan(&id,&goodsId,&userId,&avatar,&nickname,&customerUserId,&customerAvatar,&customerNickname,&question,&answer,&status,&created);err != nil {
		return
	}
	dialog = &Dialog{
		id,
		goodsId,
		userId,
		avatar,
		nickname,
		customerUserId,
		customerAvatar,
		customerNickname,
		question,
		answer,
		status,
		created,
	}
	return
}