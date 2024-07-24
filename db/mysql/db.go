package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	*Executor

	// 参数配置
	config *Config
}

// initConfig 初始化配置
func (d *Driver) initConfig() {
	d.config = &Config{
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
func (d *Driver) SetConfig(config map[string]any) {
	if d.config == nil {
		d.initConfig()
	}

	for key, value := range config {
		switch key {
		case "host":
			switch t := value.(type) {
			case string:
				d.config.host = t
			}
		case "port":
			switch t := value.(type) {
			case int:
				if t > 0 && t < 65535 {
					d.config.port = t
				}
			}
		case "username":
			switch t := value.(type) {
			case string:
				d.config.username = t
			}
		case "password":
			switch t := value.(type) {
			case string:
				d.config.password = t
			}
		case "name":
			switch t := value.(type) {
			case string:
				d.config.name = t
			}
		case "maxOpenConns":
			switch t := value.(type) {
			case int:
				d.config.maxOpenConns = t
			}
		case "maxIdleConns":
			switch t := value.(type) {
			case int:
				d.config.maxIdleConns = t
			}
		case "connMaxLifetime":
			switch t := value.(type) {
			case time.Duration:
				d.config.connMaxLifetime = t
			case int:
				d.config.connMaxLifetime = time.Duration(t) * time.Second
			case string:
				du, err := time.ParseDuration(t)
				if err == nil {
					d.config.connMaxLifetime = du
				}
			}
		}
	}
}

// GetConfig 参数配置
func (d *Driver) GetConfig() *Config {
	if d.config == nil {
		d.initConfig()
	}

	return d.config
}

// Init 初始化
func (d *Driver) Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.config.username, d.config.password, d.config.host, d.config.port, d.config.name)
	instance, err := sql.Open(d.config.driver, dsn)
	if err != nil {
		return err
	}
	instance.SetMaxOpenConns(d.config.maxOpenConns)
	instance.SetMaxIdleConns(d.config.maxIdleConns)
	instance.SetConnMaxLifetime(d.config.connMaxLifetime)

	if err := instance.Ping(); err != nil {
		return err
	}

	executor := new(Executor)
	executor.init(ExecutorTypeDb, instance, nil)
	d.Executor = executor

	return nil
}

// Tx 开启事务
func (d *Driver) Tx() (*Executor, error) {

	tx, err := d.Executor.getDb().Begin()
	if err != nil {
		return nil, err
	}

	executor := new(Executor)
	executor.init(ExecutorTypeTx, nil, tx)

	return executor, nil
}
