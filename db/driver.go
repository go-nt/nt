package db

import (
	"database/sql"
)

type Driver interface {

	// Config 参数配置
	Config(config map[string]any)

	// Init 初始化
	Init() error

	// GetValue 查询一个字段的值
	GetValue(sql string, args ...any) (string, error)

	// GetValues 查询一个字段的值
	GetValues(sql string, args ...any) ([]string, error)

	// GetMap 查询一行记录
	GetMap(sql string, args ...any) (map[string]string, error)

	// GetMaps GetRows 查询多行记录
	GetMaps(sql string, args ...any) ([]map[string]string, error)

	// Query 查询，返回查询结果集，用于 select
	Query(sql string, args ...any) (*sql.Rows, error)

	// Exec 执行，用于 insert / update / delete
	Exec(sql string, args ...any) (sql.Result, error)

	// Insert 插入数据
	Insert(table string, data map[string]any) (sql.Result, error)

	// Update 更新数据
	Update(table string, data map[string]any, primaryKeys ...string) (sql.Result, error)

	// Delete 删除
	Delete(table string, where map[string]any) (sql.Result, error)

	// Truncate 清空表
	Truncate(table string) (sql.Result, error)
}
