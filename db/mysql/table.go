package mysql

import (
	_ "github.com/go-sql-driver/mysql"
)

type Table struct {
	// 数据库
	db *Db

	// 表名
	name string

	tStruct any
}

// SetDb 设置数据库
func (table *Table) SetDb(db *Db) {
	table.db = db
}

// SetName 设置名称
func (table *Table) SetName(name string) {
	table.name = name
}
