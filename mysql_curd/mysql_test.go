package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type User struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

var db *sql.DB

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

func Test_MySQLInsert(t *testing.T) {
	if err := initMySQL(); err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	sqlStr := "INSERT INTO user(name, age) values (?,?)"
	result, err := db.Exec(sqlStr, "jpc", 18)
	if err != nil {
		fmt.Println("insert mysql failed,", err)
		return
	}
	rowsNum, _ := result.RowsAffected()
	id, _ := result.LastInsertId()
	fmt.Println(rowsNum, id)
	//time.Sleep(10 * time.Second)
}

func Test_MySQLUpdate(t *testing.T) {
	if err := initMySQL(); err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	sqlStr := "UPDATE user SET age=? where id = ?"
	result, err := db.Exec(sqlStr, "20", 2)
	if err != nil {
		fmt.Println("update mysql failed,", err)
		return
	}
	rowsNum, _ := result.RowsAffected()
	id, _ := result.LastInsertId()
	fmt.Println(rowsNum, id)
}

func Test_MySQLDelete(t *testing.T) {
	if err := initMySQL(); err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	sqlStr := "DELETE FROM user WHERE id = ?"
	result, err := db.Exec(sqlStr, 2)
	if err != nil {
		fmt.Println("delete mysql failed,", err)
		return
	}
	rowsNum, _ := result.RowsAffected()
	id, _ := result.LastInsertId()
	fmt.Println(rowsNum, id)
}

func initMySQL() (err error) {
	dsn := "root:jHw^jqx5^GWv0748@tcp(172.20.2.100:3306)/go_simple_demo?charset=utf8mb4&parseTime=True"

	if db, err = sql.Open("mysql", dsn); err != nil {
		fmt.Println("open mysql failed", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ping mysql failed", err)
		return
	}
	//最大连接数
	db.SetMaxOpenConns(10)
	//最大限制连接数
	db.SetMaxIdleConns(10)
	fmt.Println("数据库连接成功~~")
	return
}
