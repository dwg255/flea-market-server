package userModel

import (
	"database/sql"
	"flea-market/model"
	"time"
)

type User struct {
	UserId    int `json:"user_id"`
	Nickname  string `json:"nickname"`
	Openid    string `json:"openid"`
	Gender    int `json:"gender"`
	AvatarUrl string `json:"avatar_url"`
	Country   string `json:"country"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Address   string `json:"address"`
	Created   int `json:"created"`
	Phone 	  string `json:"phone"`
}

// 添加用户
func AddUser( u *User)( *User, error) {
	sqlStr := `insert into f_user (nickname, open_id, gender, avatar_url, country, province, city, address, created,phone) values(?,?,?,?,?,?,?,?,?,?)`
	if stmt, err := model.Db.Prepare(sqlStr); err != nil {
		return nil,err
	} else {
		defer stmt.Close()
		res, err := stmt.Exec(u.Nickname, u.Openid, u.Gender, u.AvatarUrl, u.Country, u.Province, u.City, u.Address, time.Now().Unix(),u.Phone)
		lastInsertId,_ := res.LastInsertId()
		u.UserId = int(lastInsertId)
		return u,err
	}
}

// 根据主键查找
func GetUserById(userId int) (u *User, err error) {
	sqlStr := `select * from f_user where user_id = ?`
	var stmt *sql.Stmt
	stmt, err =  model.Db.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId)
	u,err = ScanRow(row)
	return

}

// 根据openid查找
func GetUserByOpenId(openId string) (u *User, err error) {
	sql := `select * from f_user where open_id = ?`
	stmt, err :=  model.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(openId)
	u,err = ScanRow(row)
	return

}

func ScanRow (row *sql.Row) (u *User,err error) {
	var UserId    int
	var Nickname  string
	var Openid    string
	var Gender    int
	var AvatarUrl string
	var Country   string
	var Province  string
	var City      string
	var Address   string
	var Created   int
	var Phone string
	if err = row.Scan(&UserId, &Nickname, &Openid, &Gender, &AvatarUrl, &Country, &Province, &City, &Address, &Created,&Phone);err != nil {
		return
	}
	u = &User{
		UserId,
		Nickname,
		Openid,
		Gender,
		AvatarUrl,
		Country,
		Province,
		City,
		Address,
		Created,
		Phone,
	}
	return
}