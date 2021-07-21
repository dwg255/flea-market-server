package model

import (
	"database/sql"
	"time"
)

type User struct {
	UserId    int
	Nickname  string
	Openid    string
	Gender    int
	AvatarUrl string
	Country   string
	Province  string
	City      string
	Address   string
	Created   int
}

// 添加用户
func AddUser(db *sql.DB, u *User) error {
	sql := `insert into f_user (nickname, open_id, gender, avatar_url, country, province, city, address, created) values(?,?,?,?,?,?,?,?,?)`
	if stmt, err := db.Prepare(sql); err != nil {
		return err
	} else {
		defer stmt.Close()
		if err != nil {
			return err
		}
		_, err = stmt.Exec(u.Nickname, u.Openid, u.Gender, u.AvatarUrl, u.Country, u.Province, u.City, u.Address, time.Now().Unix())
		return nil
	}
}

// 根据主键查找
func GetUserById(userId int) (u *User, err error) {
	sql := `select * from f_user where user_id = ?`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId)
	row.Scan(u.UserId, u.Nickname, u.Openid, u.Gender, u.AvatarUrl, u.Country, u.Province, u.City, u.Address, u.Created)
	return

}

// 根据openid查找
func GetUserByOpenId(openId string) (u *User, err error) {
	sql := `select * from f_user where open_id = ?`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(openId)
	row.Scan(u.UserId, u.Nickname, u.Openid, u.Gender, u.AvatarUrl, u.Country, u.Province, u.City, u.Address, u.Created)
	return

}
