package db

import (
	"database/sql"
	"fmt"
	"parser/loger"
	"parser/telegram"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func mysqlConn(dbName string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("root:Passwd@unix(/var/run/mysql.sock)/%s", dbName))
	if err != nil {
		panic(err.Error())
	}
	return db
}

func selectHash(db *sql.DB, hash string) bool {
	var rowHash int
	db.QueryRow("select id from golang_python where link_hash = ?", hash).Scan(&rowHash)
	if rowHash != 0 {
		return true
	}
	return false
}

func insertHash(db *sql.DB, url, page, text, hash string, times int64) {
	resu, err := db.Prepare("INSERT INTO golang_python (site, page_link, page_text, timestamp, link_hash)")
	if err != nil {
		loger.ForError(err)
	}

	if _, err = resu.Exec(url, page, text, times, hash); err != nil {
		loger.ForError(err)
	}
}

func CheckSiteNewBot(url, page, text, hash string) {
	dataBase := mysqlConn("news-bot")
	checkLink := selectHash(dataBase, hash)
	times := time.Now().Unix()
	if checkLink == false {
		insertHash(dataBase, url, page, text, hash, times)
		telegram.SendMessage(text)
	}
	defer dataBase.Close()
}
