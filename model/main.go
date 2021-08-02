package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriverName = "sqlite3"
	dbName       = "./db/data.db3"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open(dbDriverName, dbName)
	if checkErr(err) {
		return
	}
	err = createTable(Db)
	checkErr(err)
}

func createTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS "f_user" (
		"user_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"nickname" TEXT(32),
		"open_id" text(32) NOT NULL,
		"gender" integer,
		"avatar_url" TEXT,
		"country" TEXT,
		"province" TEXT,
		"city" TEXT,
		"address" TEXT,
		"created" integer
	  )`
	if _, err := db.Exec(sql); err != nil {
		return err
	}

	sql = `CREATE TABLE IF NOT EXISTS "f_goods" (
		"goods_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"price" REAL,
		"pics" TEXT,
		"user_id" INTEGER,
		"nickname" TEXT,
		"avatar_url" TEXT,
		"address" TEXT,
		"latitude" real,
		"longitude" real,
		"tags" TEXT,
		"content" TEXT,
		"new_message" TEXT,
		"goods_num" INTEGER,
		"star_num" INTEGER,
		"fav_num" INTEGER,
		"views_num" INTEGER,
		"created" INTEGER
        "online_sell"  INTEGER,
		"express_type"  INTEGER,
    	"cat_id" INTEGER,
    	"status" INTEGER                    
	  );`
	if _, err := db.Exec(sql); err != nil {
		return err
	}
	sql = `CREATE TABLE IF NOT EXISTS  "f_dialog" (
		"id"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		"goods_id"  INTEGER,
		"user_id"  INTEGER,
		"avatar"  TEXT,
		"nickname"  TEXT,
		"customer_user_id"  INTEGER,
		"customer_avatar"  TEXT,
		"customer_nickname"  TEXT,
		"question"  TEXT,
		"answer"  TEXT,
		"status"  INTEGER,
		"created"  INTEGER
		);`
	if _, err := db.Exec(sql); err != nil {
		return err
	}
	return nil
}
func checkErr(e error) bool {
	if e != nil {
		log.Fatal(e)
		fmt.Println("fatal err",e.Error())
		return true
	}
	return false
}
