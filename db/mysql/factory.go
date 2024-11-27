package mysql

import (
	"errors"
	"time"

	"github.com/go-ini/ini"
)

type Config struct {

	// 驱动：mysql | sqllite3 | postgress
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

var configs map[string]*Config
var drivers map[string]*Driver

// initConfig 初始化配置
func initConfig() *Config {
	return &Config{
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

// SetIniConfig 配置
func SetConfig(name string, c map[string]any) error {
	if configs == nil {
		configs = make(map[string]*Config)
	}

	config := initConfig()

	for key, value := range c {
		switch key {
		case "host":
			switch t := value.(type) {
			case string:
				config.host = t
			}
		case "port":
			switch t := value.(type) {
			case int:
				if t > 0 && t < 65535 {
					config.port = t
				} else {
					return errors.New("mysql config parameter(port) is not a valid value")
				}
			}
		case "username":
			switch t := value.(type) {
			case string:
				config.username = t
			}
		case "password":
			switch t := value.(type) {
			case string:
				config.password = t
			}
		case "name":
			switch t := value.(type) {
			case string:
				config.name = t
			}
		case "maxOpenConns":
			switch t := value.(type) {
			case int:
				config.maxOpenConns = t
			}
		case "maxIdleConns":
			switch t := value.(type) {
			case int:
				config.maxIdleConns = t
			}
		case "connMaxLifetime":
			switch t := value.(type) {
			case time.Duration:
				config.connMaxLifetime = t
			case int:
				config.connMaxLifetime = time.Duration(t) * time.Second
			case string:
				du, err := time.ParseDuration(t)
				if err == nil {
					config.connMaxLifetime = du
				}
			}
		}
	}

	configs[name] = config

	return nil
}

// SetIniConfig 设置 ini 配置
func SetIniConfig(name string, section *ini.Section) error {
	if configs == nil {
		configs = make(map[string]*Config)
	}

	config := initConfig()

	section.MapTo(config)

	if config.host == "" {
		return errors.New("mysql config parameter(host) not found in ini section")
	}

	if config.port <= 0 || config.port >= 65535 {
		return errors.New("mysql config parameter(port) is not a valid value")
	}

	configs[name] = config

	return nil
}

// GetConfigs 获取配置项
func GetConfigs(name string) map[string]*Config {
	return configs
}

// GetConfig 获取配置项
func GetConfig(name string) (*Config, error) {
	config, ok := configs[name]
	if ok {
		return config, nil
	}

	return nil, errors.New("mysql config (" + name + ") not found")
}

// GetDb 获取数据库实例
func GetDb(name string) (*Driver, error) {
	d, ok := drivers[name]
	if ok {
		return d, nil
	}

	config, ok := configs[name]
	if ok {
		d := new(Driver)
		d.SetConfig(config)
		err := d.Init()
		if err != nil {
			return nil, err
		}

		if drivers == nil {
			drivers = make(map[string]*Driver)
		}

		drivers[name] = d
		return d, nil
	}

	return nil, errors.New("mysql (" + name + ") not found")
}
