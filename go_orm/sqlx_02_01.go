package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
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
	queryEmployees()
}

func queryEmployees() {
	sqlStr := "SELECT id, name, department, salary FROM employees"
	var employees []Employees
	if err := db.Select(&employees, sqlStr); err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return
	}

	for i := 0; i < len(employees); i++ {
		fmt.Printf("%d, %s, %s, %s\n", employees[i].Id, employees[i].Name, employees[i].Department, employees[i].Salary)
	}
}

type Employees struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     string `db:"salary"`
}
