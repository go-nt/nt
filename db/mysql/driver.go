package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Driver struct {
	*Executor

	// 参数配置
	config *Config
}

// SetConfig 参数配置
func (d *Driver) SetConfig(config *Config) {
	d.config = config
}

// GetConfig 参数配置
func (d *Driver) GetConfig() *Config {
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
