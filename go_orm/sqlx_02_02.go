package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 题目2：实现类型安全映射
//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
//要求 ：
//定义一个 Book 结构体，包含与 books 表对应的字段。
//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

var db *sqlx.DB // 定义全局变量db

func initMySQL() (err error) {
	dsn := "test:123456@tcp(127.0.0.1:3306)/test"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return
	}
	// 测试连接
	err = db.Ping()
	if err != nil {
		fmt.Printf("ping database failed, err:%v\n", err)
		return
	}
	// 设置连接池参数
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}
func main() {
	// 初始化数据库
	err := initMySQL()
	if err != nil {
		fmt.Printf("init MySQL failed, err:%v\n", err)
		return
	}
	fmt.Println("MySQL initialized successfully")
	queryBooks()
}

func queryBooks() {
	sqlStr := "SELECT id , title , author ,price FROM books WHERE price > ?"
	var books []Books
	if err := db.Select(&books, sqlStr, 50); err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return
	}

	for i := 0; i < len(books); i++ {
		fmt.Printf("%d, %s, %s, %.2f\n", books[i].Id, books[i].Title, books[i].Author, books[i].Price)
	}
}

type Books struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}
