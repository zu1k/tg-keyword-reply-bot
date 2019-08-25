package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db = &sql.DB{}


func dbinit(token string) {
	_dbcreate()
	stmt, err := db.Prepare("INSERT INTO settings(key, value) values(?,?)")
	checkErr(err)
	_, err = stmt.Exec("token", token)
	checkErr(err)
}


func dbaddgroup(groupid int64) {
	stmt, err := db.Prepare("INSERT INTO keyword_reply(groupid, kvjson) values(?,?)")
	checkErr(err)
	_, err = stmt.Exec(groupid, "")
	checkErr(err)
}


func dbupdategroup(groupid int64, kvjson string) {
	stmt, err := db.Prepare("update keyword_reply set kvjson=? where groupid=?")
	checkErr(err)
	_, err = stmt.Exec(kvjson, groupid)
	checkErr(err)
}


func dbread() {
	rows, err := db.Query("SELECT * FROM keyword_reply")
	checkErr(err)
	var groupid int64
	var kvjson string

	for rows.Next() {
		err = rows.Scan(&groupid, &kvjson)
		checkErr(err)
		kvs := json2kvs(kvjson)
		allkvs[groupid] = kvs
		groups = append(groups, groupid)
	}
	rows.Close()

	rows, err = db.Query("SELECT value FROM settings")
	checkErr(err)
	var value string
	rows.Next()
	err = rows.Scan(&value)
	checkErr(err)
	token = value
	rows.Close()
}

func dbopen() {
	dbb, err := sql.Open("sqlite3", "./bot.db")
	checkErr(err)
	db = dbb
}

func dbclose() {
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


func _dbcreate() {
	settings := `
    CREATE TABLE IF NOT EXISTS settings(
        key VARCHAR(64) PRIMARY KEY,
        value VARCHAR(200) NULL 
    );
    `

	kwr := `
    CREATE TABLE IF NOT EXISTS keyword_reply(
		groupid INTEGER PRIMARY KEY,
		kvjson VARCHAR(500000) NULL		
    );
    `

	db.Exec(settings)
	db.Exec(kwr)
}
