package mysql

import (
	_ "github.com/go-sql-driver/mysql"
)

type Tuple struct {
	// 数据库
	db *Db

	// 表名
	name string

	tStruct any
}

// SetDb 设置数据库
func (tuple *Tuple) SetDb(db *Db) {
	tuple.db = db
}

// SetName 设置名称
func (tuple *Tuple) SetName(name string) {
	tuple.name = name
}

// SetStruct 设置结构
func (tuple *Tuple) SetStruct(s any) {
	tuple.tStruct = s
}
