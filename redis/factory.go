package redis

import (
	"errors"

	"github.com/go-ini/ini"
)

var configs map[string]map[string]any
var drivers map[string]*Driver

func SetConfig(name string, config map[string]any) {
	if configs == nil {
		configs = make(map[string]map[string]any)
	}
	configs[name] = config
}

// FormatIniConfig 格式化 ini 配置
func FormatIniConfig(section *ini.Section) (map[string]any, error) {

	configKeyHostString := "127.0.0.1"
	configKeyHost, err := section.GetKey("host")
	if err == nil {
		configKeyHostString = configKeyHost.String()
		if err != nil || configKeyHostString == "" {
			return nil, errors.New("redis config parameter(host) is not a valid value")
		}
	}

	configKeyPortInt := 6379
	configKeyPort, err := section.GetKey("port")
	if err == nil {
		configKeyPortInt, err = configKeyPort.Int()
		if err != nil || configKeyPortInt <= 0 || configKeyPortInt >= 65535 {
			return nil, errors.New("redis config parameter(port) is not a valid value")
		}
	}

	configKeyPasswordString := ""
	configKeyPassword, err := section.GetKey("host")
	if err == nil {
		configKeyPasswordString = configKeyPassword.String()
	}

	configKeyDbInt := 0
	configKeyDb, err := section.GetKey("db")
	if err == nil {
		configKeyDbInt, err = configKeyDb.Int()
		if err != nil || configKeyDbInt <= 0 || configKeyDbInt >= 65535 {
			return nil, errors.New("redis config parameter(db) is not a valid value")
		}
	}

	return map[string]any{
		"host":     configKeyHostString,
		"port":     configKeyPortInt,
		"password": configKeyPasswordString,
		"db":       configKeyDbInt,
	}, nil
}

// GetConfigs 获取配置项
func GetConfigs(name string) map[string]map[string]any {
	return configs
}

// GetConfig 获取配置项
func GetConfig(name string) (map[string]any, error) {
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
		d, err := GetRedisByConfig(config)
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

// GetRedisByConfig 按配置文件创建Redis实例
func GetRedisByConfig(config map[string]any) (*Driver, error) {
	d := new(Driver)
	d.SetConfig(config)
	err := d.Init()
	if err != nil {
		return nil, err
	}

	return d, nil
}
