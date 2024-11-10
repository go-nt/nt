package mysql

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type ExecutorType string

const (
	ExecutorTypeDb ExecutorType = "db"
	ExecutorTypeTx ExecutorType = "tx"
)

type Executor struct {
	executorType ExecutorType
	db           *sql.DB
	tx           *sql.Tx
}

// init 初始化
func (e *Executor) init(executorType ExecutorType, db *sql.DB, tx *sql.Tx) {
	e.executorType = executorType
	e.db = db
	e.tx = tx
}

// getDb 数据库连接对象
func (e *Executor) getDb() *sql.DB {
	return e.db
}

// getTx 数据库连接对象
func (e *Executor) getTx() *sql.Tx {
	return e.tx
}

// GetTable 获取表记录
func (e *Executor) GetTable(table string) *Table {
	t := new(Table)
	t.SetExecutor(e)
	t.SetName(table)
	t.Init()
	return t
}

// GetTuple 获取行记录
func (e *Executor) GetTuple(table string) *Tuple {
	t := new(Tuple)
	t.SetExecutor(e)
	t.SetName(table)
	t.Init()
	return t
}

// GetValue 查询一个字段的值
func (e *Executor) GetValue(sq string, args ...any) (string, error) {
	var rows *sql.Rows
	var err error
	if e.executorType == ExecutorTypeDb {
		rows, err = e.db.Query(sq, args...)
	} else {
		rows, err = e.tx.Query(sq, args...)
	}
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		var val string
		rows.Scan(&val)
		return val, nil
	}
	return "", errors.New("db->GetValue no matched result")
}

// GetValues 查询一个字段的值
func (e *Executor) GetValues(sq string, args ...any) ([]string, error) {
	var rows *sql.Rows
	var err error
	if e.executorType == ExecutorTypeDb {
		rows, err = e.db.Query(sq, args...)
	} else {
		rows, err = e.tx.Query(sq, args...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []string
	if rows.Next() {
		var val string
		if err = rows.Scan(&val); err != nil {
			return nil, err
		}
		values = append(values, val)
	}
	return values, nil
}

// GetMap 查询一行记录
func (e *Executor) GetMap(sq string, args ...any) (map[string]string, error) {
	var rows *sql.Rows
	var err error
	if e.executorType == ExecutorTypeDb {
		rows, err = e.db.Query(sq, args...)
	} else {
		rows, err = e.tx.Query(sq, args...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		columns, _ := rows.Columns()
		columnLen := len(columns)

		columnData := make([]string, columnLen)
		columnDataPointers := make([]any, columnLen)
		for i := 0; i < columnLen; i = i + 1 {
			columnDataPointers[i] = &columnData[i]
		}

		if err = rows.Scan(columnDataPointers...); err != nil {
			return nil, err
		}

		m := make(map[string]string)
		for i, colName := range columns {
			val := columnDataPointers[i].(*string)
			m[colName] = *val
		}

		return m, nil
	}

	return nil, errors.New("db->GetMap no matched results")
}

// GetBind 查询记录, 缓定到指定对象
func (e *Executor) GetBind(bind *any, sq string, args ...any) error {
	// TODO
	return nil
}

// GetMaps 查询多行记录
func (e *Executor) GetMaps(sq string, args ...any) ([]map[string]string, error) {
	var rows *sql.Rows
	var err error
	if e.executorType == ExecutorTypeDb {
		rows, err = e.db.Query(sq, args...)
	} else {
		rows, err = e.tx.Query(sq, args...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	columnLen := len(columns)

	var maps []map[string]string
	for rows.Next() {
		columnData := make([]string, columnLen)
		columnDataPointers := make([]any, columnLen)
		for i := 0; i < columnLen; i = i + 1 {
			columnDataPointers[i] = &columnData[i]
		}

		if err = rows.Scan(columnDataPointers...); err != nil {
			return nil, err
		}

		m := make(map[string]string)
		for i, colName := range columns {
			val := columnDataPointers[i].(*string)
			m[colName] = *val
		}

		maps = append(maps, m)
	}
	return maps, nil
}

// Query 查询，返回查询结果集，用于 select
func (e *Executor) Query(sq string, args ...any) (*sql.Rows, error) {
	if e.executorType == ExecutorTypeDb {
		return e.db.Query(sq, args...)
	} else {
		return e.tx.Query(sq, args...)
	}
}

// Exec 执行，用于 insert / update / delete
func (e *Executor) Exec(sq string, args ...any) (sql.Result, error) {
	if e.executorType == ExecutorTypeDb {
		return e.db.Exec(sq, args...)
	} else {
		return e.tx.Exec(sq, args...)
	}
}

// Insert 插入数据
func (e *Executor) Insert(table string, data map[string]any) (sql.Result, error) {
	sq := "INSERT INTO " + table + "("
	vs := ""
	var args []any

	isFirst := false
	for k, v := range data {
		if !isFirst {
			isFirst = true
		} else {
			sq += ","
			vs += ","
		}

		sq += k
		vs += "?"

		args = append(args, v)
	}

	sq += ") VALUES (" + vs + ")"

	if e.executorType == ExecutorTypeDb {
		return e.db.Exec(sq, args...)
	} else {
		return e.tx.Exec(sq, args...)
	}
}

// Update 更新数据
func (e *Executor) Update(table string, data map[string]any, primaryKeys ...string) (sql.Result, error) {

	sq := "UPDATE " + table + " SET "

	where := make(map[string]any)

	var args []any

	isFirst := false
	for k, v := range data {

		for _, primaryKey := range primaryKeys {
			if k == primaryKey {
				where[k] = v
				goto nextData
			}
		}

		if !isFirst {
			isFirst = true
		} else {
			sq += ","
		}

		sq += k + "=?"

		args = append(args, v)

	nextData:
	}

	if len(where) > 0 {
		sq += " WHERE "
		isFirst = false
		for k, v := range where {
			if !isFirst {
				isFirst = true
			} else {
				sq += ","
			}

			sq += k + "=?"

			args = append(args, v)
		}
	}

	if e.executorType == ExecutorTypeDb {
		return e.db.Exec(sq, args...)
	} else {
		return e.tx.Exec(sq, args...)
	}
}

// Delete 删除
func (e *Executor) Delete(table string, where map[string]any) (sql.Result, error) {
	sq := "DELETE FROM " + table + " WHERE "

	var args []any
	isFirst := false
	for k, v := range where {
		if !isFirst {
			isFirst = true
		} else {
			sq += " AND "
		}

		sq += k + "=?"

		args = append(args, v)
	}

	if e.executorType == ExecutorTypeDb {
		return e.db.Exec(sq, args...)
	} else {
		return e.tx.Exec(sq, args...)
	}
}

// Truncate 清空表
func (e *Executor) Truncate(table string) (sql.Result, error) {
	sq := "TRUNCATE " + table

	if e.executorType == ExecutorTypeDb {
		return e.db.Exec(sq)
	} else {
		return e.tx.Exec(sq)
	}
}
