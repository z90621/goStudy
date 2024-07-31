package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id        int
	name      string
	age       int
	className string
}

type DbWorker struct {
	Dsn string
	Db  *sql.DB
}

func main() {
	dbw := DbWorker{Dsn: "root:root@tcp(localhost:3306)/sat?charset=utf8mb4"}
	db, err := sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	dbw.Db = db

	// stu := Student{id: 1, name: "zs", age: 18, className: "math"}
	// dbw.insert(stu)
	dbw.selectAll()
}

func (db *DbWorker) insert(stu Student) Student {
	stmt, _ := db.Db.Prepare(`INSERT INTO STUDENT(ID,NAME,AGE,CLASSNAME) VALUES (?,?,?,?)`)
	defer stmt.Close()
	ret, err := stmt.Exec(stu.id, stu.name, stu.age, stu.className)
	if err != nil {
		fmt.Println(err)
	}
	if lastId, err := ret.LastInsertId(); err == nil {
		fmt.Println("lastid->", lastId)
	}
	if aff, err := ret.RowsAffected(); err == nil {
		fmt.Println("affectRow->", aff)
	}
	return stu

}

func (db *DbWorker) selectAll() []Student {
	// 如果方法包含Query，那么这个方法是用于查询并返回rows的。其他用Exec()
	var stus []Student
	rows, _ := db.Db.Query(`SELECT * FROM STUDENT WHERE NAME = ?`, "zs")
	//err = db.QueryRow("select name from users where id = ?", 1).Scan(&name) // 单行查询，直接处理
	defer rows.Close()

	columns, _ := rows.Columns()
	fmt.Println(columns)

	for rows.Next() {
		var s Student
		err := rows.Scan(&s.id, &s.name, &s.age, &s.className)
		if err != nil {
			log.Fatal(err)
		}
		stus = append(stus, s)
	}
	for _, v := range stus {
		fmt.Println(v)
	}
	return stus
}

// func (db *DbWorker) openTransaction() {
// 	tx,err  := db.Db.Begin()
// 	tx.Commit()
// 	tx.Rollback()
// }
