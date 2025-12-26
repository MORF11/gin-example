package tls

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

type User struct {
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

var db *sql.DB
var e error

func DBClose() {
	db.Close()
}

func Delu(id string) (int64,error) {
	res,_ :=db.Exec("DELETE FROM users WHERE id = ?",id)
	return res.RowsAffected()
}

func Addu(nm string) {
	db.Exec("INSERT INTO users (name) VALUES (?)",nm)
}

func Retu() []User {
	var res[]User
	dt,e := db.Query("SELECT * FROM users")
	if e != nil {
		fmt.Println("error query",e)
		return make([]User, 0)
	}
	for dt.Next() {
		var u User
		if e := dt.Scan(&u.Id,&u.Name); e != nil {
			fmt.Println("error scan",e)
			return make([]User, 0)
		}
		res = append(res, u)
	}
	dt.Close()
	return res
}

func Check_dir() []string {
	q,e := os.ReadDir("./files")
	var res []string
	if e != nil {
		return make([]string,0)
	}
	for _,f := range q {
		res = append(res, f.Name())
	}
	return res
}

func init() {
	db,e = sql.Open("sqlite","./db.db")
	if e != nil {
		fmt.Println("db init error",e)
		return
	}
	db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT NOT NULL)")
}