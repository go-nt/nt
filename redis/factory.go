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
var drivers map[string]*Driver

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

// FormatIniConfig 格式化 ini 配置
func SetIniConfig(name string, section *ini.Section) error {
	if configs == nil {
		configs = make(map[string]*Config)
	}

	config := initConfig()

	configKeyHost, err := section.GetKey("host")
	if err == nil {
		t := configKeyHost.String()
		if t != "" {
			config.host = t
		} else {
			return errors.New("redis config parameter(host) is not a valid value")
		}
	}

	configKeyPort, err := section.GetKey("port")
	if err == nil {
		t, err := configKeyPort.Int()
		if err == nil && t > 0 && t < 65535 {
			config.port = t
		} else {
			return errors.New("redis config parameter(port) is not a valid value")
		}
	}

	configKeyPassword, err := section.GetKey("password")
	if err == nil {
		t := configKeyPassword.String()
		config.password = t
	}

	configKeyDb, err := section.GetKey("db")
	if err == nil {
		t, err := configKeyDb.Int()
		if err == nil && t >= 0 && t <= 65535 {
			config.port = t
		} else {
			return errors.New("redis config parameter(db) is not a valid value")
		}
	}

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

	return nil, errors.New("redis config (" + name + ") not found")
}

// GetRedis 获取Redis实例
func GetRedis(name string) (*Driver, error) {
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

	return nil, errors.New("redis (" + name + ") not found")
}
