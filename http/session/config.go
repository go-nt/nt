package session

import (
	"errors"

	"github.com/go-ini/ini"
)

var config *Config

type Config struct {
	// 名乐
	name string

	// 超时时间
	expire int

	// 驱动
	driver string

	// 区动为 redis 时， 指定 redis name
	redis string
}

// initConfig 初始化配置
func initConfig() {
	config = &Config{
		name:   "SSID",
		expire: 1440,
		driver: "file",
		redis:  "",
	}
}

// SetConfig 参数配置
func SetConfig(c map[string]any) error {
	if config == nil {
		initConfig()
	}

	for key, value := range c {
		switch key {
		case "name":
			switch t := value.(type) {
			case string:
				if t != "" {
					config.name = t
				} else {
					return errors.New("session config parameter(name) is not a valid value")
				}
			}
		case "expire":
			switch t := value.(type) {
			case int:
				if t > 0 {
					config.expire = t
				} else {
					return errors.New("session config parameter(expire) is not a valid value")
				}
			}
		case "driver":
			switch t := value.(type) {
			case string:
				if t == "file" || t == "redis" {
					config.driver = t
				} else {
					return errors.New("session config parameter(driver) is not a valid value")
				}
			}
		case "redis":
			switch t := value.(type) {
			case string:
				config.redis = t
			}
		}
	}

	return nil
}

// FormatIniConfig 格式化 ini 配置
func SetIniConfig(section *ini.Section) error {
	if config == nil {
		initConfig()
	}

	section.MapTo(config)

	if config.name == "" {
		return errors.New("session config parameter(name) is not a valid value")
	}

	if config.expire <= 0 {
		return errors.New("session config parameter(expire) is not a valid value")
	}

	if config.driver != "file" && config.driver != "redis" {
		return errors.New("session config parameter(driver) is not a valid value")
	}

	if config.driver == "redis" {
		if config.redis == "" {
			return errors.New("session config parameter(driver) is not a valid value")
		}
	}

	return nil
}
