package mysql

import (
	_ "github.com/go-sql-driver/mysql"
)

type Tuple struct {
	// 搪行器
	executor *Executor

	// 表名
	name string

	tStruct any
}

// Init 初始化
func (tuple *Tuple) Init() *Tuple {
	return tuple
}

// SetExecutor 设置执行品
func (tuple *Tuple) SetExecutor(e *Executor) *Tuple {
	tuple.executor = e
	return tuple
}

// SetName 设置名称
func (tuple *Tuple) SetName(name string) {
	tuple.name = name
}

// SetStruct 设置结构
func (tuple *Tuple) SetStruct(s any) {
	tuple.tStruct = s
}
