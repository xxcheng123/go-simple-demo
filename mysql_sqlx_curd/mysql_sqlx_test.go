package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"testing"
)

type User struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

var db *sqlx.DB

func Test_MySQLQueryRow(t *testing.T) {
	if err := initMySQL(); err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	sqlStr := "SELECT id,name,age from user where id=?"
	var u User
	if err := db.QueryRow(sqlStr, "1").Scan(&u.ID, &u.Name, &u.Age); err != nil {
		fmt.Println("query mysql failed,", err)
		return
	}
	fmt.Println(u)
	//time.Sleep(10 * time.Second)
}
func Test_MySQLGet(t *testing.T) {
	if err := initMySQL(); err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	sqlStr := "SELECT id,name,age from user where id>=?"
	var u User
	if err := db.Get(&u, sqlStr, "1"); err != nil {
		fmt.Println("query mysql failed,", err)
		return
	}
	fmt.Println(u)
	//time.Sleep(10 * time.Second)
}
func Test_MySQLSelect(t *testing.T) {
	if err := initMySQL(); err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	sqlStr := "SELECT id,name,age from user where id>=?"
	us := new([]User)
	if err := db.Select(us, sqlStr, "1"); err != nil {
		fmt.Println("query mysql failed,", err)
		return
	}
	fmt.Println(us)
	//time.Sleep(10 * time.Second)
}

func initMySQL() (err error) {
	dsn := "root:jHw^jqx5^GWv0748@tcp(172.20.2.100:3306)/go_simple_demo?charset=utf8mb4&parseTime=True"
	//自带PING
	if db, err = sqlx.Connect("mysql", dsn); err != nil {
		fmt.Println("open mysql failed", err)
		return
	}

	//最大连接数
	db.SetMaxOpenConns(10)
	//最大限制连接数
	db.SetMaxIdleConns(10)
	fmt.Println("数据库连接成功~~")
	return
}
