package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DbConfig struct {
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

type Db struct {
	Executor

	// 参数配置
	config *DbConfig

	executor *Executor
}

// initConfig 初始化配置
func (db *Db) initConfig() {
	db.config = &DbConfig{
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

// SetConfig 参数配置
func (db *Db) SetConfig(config map[string]any) {
	if db.config == nil {
		db.initConfig()
	}

	for key, value := range config {
		switch key {
		case "host":
			switch t := value.(type) {
			case string:
				db.config.host = t
			}
		case "port":
			switch t := value.(type) {
			case int:
				if t > 0 && t < 65535 {
					db.config.port = t
				}
			}
		case "username":
			switch t := value.(type) {
			case string:
				db.config.username = t
			}
		case "password":
			switch t := value.(type) {
			case string:
				db.config.password = t
			}
		case "name":
			switch t := value.(type) {
			case string:
				db.config.name = t
			}
		case "maxOpenConns":
			switch t := value.(type) {
			case int:
				db.config.maxOpenConns = t
			}
		case "maxIdleConns":
			switch t := value.(type) {
			case int:
				db.config.maxIdleConns = t
			}
		case "connMaxLifetime":
			switch t := value.(type) {
			case time.Duration:
				db.config.connMaxLifetime = t
			case int:
				db.config.connMaxLifetime = time.Duration(t) * time.Second
			case string:
				d, err := time.ParseDuration(t)
				if err == nil {
					db.config.connMaxLifetime = d
				}
			}
		}
	}
}

// GetConfig 参数配置
func (db *Db) GetConfig() *DbConfig {
	if db.config == nil {
		db.initConfig()
	}

	return db.config
}

// Init 初始化
func (db *Db) Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.config.username, db.config.password, db.config.host, db.config.port, db.config.name)
	instance, err := sql.Open(db.config.driver, dsn)
	if err != nil {
		return err
	}
	instance.SetMaxOpenConns(db.config.maxOpenConns)
	instance.SetMaxIdleConns(db.config.maxIdleConns)
	instance.SetConnMaxLifetime(db.config.connMaxLifetime)

	if err := instance.Ping(); err != nil {
		return err
	}

	executor := new(Executor)
	executor.init(ExecutorTypeDb, instance, nil)
	db.executor = executor

	return nil
}

// Tx 开启事务
func (db *Db) Tx() (*Executor, error) {

	tx, err := db.executor.getDb().Begin()
	if err != nil {
		return nil, err
	}

	executor := new(Executor)
	executor.init(ExecutorTypeDb, nil, tx)

	return executor, nil
}
