package redis

import (
	"errors"

	"github.com/go-ini/ini"
)

type Config struct {
	// 主机名
	host string

	// 端口号
	port int

	// 密码
	password string

	// 数据库
	db int
}

var configs map[string]*Config

// initConfig 初始化配置
func initConfig() *Config {
	return &Config{
		host:     "127.0.0.1",
		port:     6379,
		password: "",
		db:       0,
	}
}

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
				if t != "" {
					config.host = t
				} else {
					return errors.New("redis config parameter(host) is not a valid value")
				}
			}
		case "port":
			switch t := value.(type) {
			case int:
				if t > 0 && t < 65535 {
					config.port = t
				} else {
					return errors.New("redis config parameter(port) is not a valid value")
				}
			}
		case "password":
			switch t := value.(type) {
			case string:
				config.password = t
			}
		case "db":
			switch t := value.(type) {
			case int:
				if t >= 0 && t <= 65535 {
					config.port = t
				} else {
					return errors.New("redis config parameter(db) is not a valid value")
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
		return errors.New("redis config parameter(host) is not a valid value")
	}

	if config.port <= 0 || config.port >= 65535 {
		return errors.New("redis config parameter(port) is not a valid value")
	}

	if config.db < 0 || config.db > 65535 {
		return errors.New("redis config parameter(db) is not a valid value")
	}

	configs[name] = config

	return nil
}
