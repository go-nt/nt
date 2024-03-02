package driver

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Config struct {
	// 驱动
	driver string

	// 主机名
	host string

	// 端口号
	port int

	// 用户名
	username string

	// 密码
	password string

	// 数据库名
	name string

	// 最长生命周期
	connMaxLifetime time.Duration

	// 最大空闲连接数
	maxIdleConns int

	// 最大连接数，0-不限制
	maxOpenConns int
}

type Driver struct {
	// 参数配置
	config *Config

	instance *sql.DB
}

// initConfig 初始化配置
func (driver *Driver) initConfig() {
	driver.config = &Config{
		driver:          "mysql",
		host:            "127.0.0.1",
		port:            3306,
		username:        "root",
		password:        "",
		name:            "go-nt",
		maxOpenConns:    4,
		maxIdleConns:    4,
		connMaxLifetime: time.Minute * 10,
	}
}

// Config 参数配置
func (driver *Driver) Config(config map[string]any) {
	if driver.config == nil {
		driver.initConfig()
	}

	for key, value := range config {
		switch key {
		case "host":
			switch t := value.(type) {
			case string:
				driver.config.host = t
			}
		case "port":
			switch t := value.(type) {
			case int:
				if t > 0 && t < 65535 {
					driver.config.port = t
				}
			}
		case "username":
			switch t := value.(type) {
			case string:
				driver.config.username = t
			}
		case "password":
			switch t := value.(type) {
			case string:
				driver.config.password = t
			}
		case "name":
			switch t := value.(type) {
			case string:
				driver.config.name = t
			}
		case "maxOpenConns":
			switch t := value.(type) {
			case int:
				driver.config.maxOpenConns = t
			}
		case "maxIdleConns":
			switch t := value.(type) {
			case int:
				driver.config.maxIdleConns = t
			}
		case "connMaxLifetime":
			switch t := value.(type) {
			case time.Duration:
				driver.config.connMaxLifetime = t
			case int:
				driver.config.connMaxLifetime = time.Duration(t) * time.Second
			case string:
				d, err := time.ParseDuration(t)
				if err == nil {
					driver.config.connMaxLifetime = d
				}
			}
		}
	}
}

// Init 初始化
func (driver *Driver) Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", driver.config.username, driver.config.password, driver.config.host, driver.config.port, driver.config.name)
	instance, err := sql.Open(driver.config.driver, dsn)
	if err != nil {
		return err
	}
	instance.SetMaxOpenConns(driver.config.maxOpenConns)
	instance.SetMaxIdleConns(driver.config.maxIdleConns)
	instance.SetConnMaxLifetime(driver.config.connMaxLifetime)

	if err := instance.Ping(); err != nil {
		return err
	}
	driver.instance = instance

	return nil
}

// GetValue 查询一个字段的值
func (driver *Driver) GetValue(sql string, args ...any) (string, error) {
	rows, err := driver.instance.Query(sql, args...)
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
func (driver *Driver) GetValues(sql string, args ...any) ([]string, error) {
	rows, err := driver.instance.Query(sql, args...)
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
func (driver *Driver) GetMap(sql string, args ...any) (map[string]string, error) {
	rows, err := driver.instance.Query(sql, args...)
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

// GetMaps 查询多行记录
func (driver *Driver) GetMaps(sql string, args ...any) ([]map[string]string, error) {

	rows, err := driver.instance.Query(sql, args...)
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
func (driver *Driver) Query(sql string, args ...any) (*sql.Rows, error) {
	return driver.instance.Query(sql, args...)
}

// Exec 执行，用于 insert / update / delete
func (driver *Driver) Exec(sql string, args ...any) (sql.Result, error) {
	return driver.instance.Exec(sql, args...)
}

// Insert 插入数据
func (driver *Driver) Insert(table string, data map[string]any) (sql.Result, error) {
	sq := "INSERT INTO " + table + "("
	vs := ""
	var args []any

	isFirst := false
	for k, v := range data {
		if isFirst == false {
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

	return driver.instance.Exec(sq, args...)
}

// Update 更新数据
func (driver *Driver) Update(table string, data map[string]any, primaryKeys ...string) (sql.Result, error) {

	sq := "UPDATE " + table + " SET "

	var where map[string]any

	var args []any

	isFirst := false
	for k, v := range data {

		for _, primaryKey := range primaryKeys {
			if k == primaryKey {
				where[k] = v
				goto nextData
			}
		}

		if isFirst == false {
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
			if isFirst == false {
				isFirst = true
			} else {
				sq += ","
			}

			sq += k + "=?"

			args = append(args, v)
		}
	}

	return driver.instance.Exec(sq, args...)
}

// Delete 删除
func (driver *Driver) Delete(table string, where map[string]any) (sql.Result, error) {
	sq := "DELETE FROM " + table + " WHERE "

	var args []any
	isFirst := false
	for k, v := range where {
		if isFirst == false {
			isFirst = true
		} else {
			sq += " AND "
		}

		sq += k + "=?"

		args = append(args, v)
	}

	return driver.instance.Exec(sq, args...)
}

// Truncate 清空表
func (driver *Driver) Truncate(table string) (sql.Result, error) {
	s := "TRUNCATE " + table
	return driver.instance.Exec(s)
}
