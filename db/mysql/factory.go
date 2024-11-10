package mysql

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
	configKeyHost, err := section.GetKey("host")
	if err != nil {
		return nil, errors.New("mysql config parameter(host) not found in ini section")
	}

	configKeyPort, err := section.GetKey("port")
	if err != nil {
		return nil, errors.New("mysql config parameter(port) not found in ini section")
	}

	configKeyPortInt, err := configKeyPort.Int()
	if err != nil || configKeyPortInt <= 0 || configKeyPortInt >= 65535 {
		return nil, errors.New("mysql config parameter(port) is not a valid value")
	}

	configKeyUsername, err := section.GetKey("username")
	if err != nil {
		return nil, errors.New("mysql config parameter(username) not found in ini section")
	}

	configKeyPassword, err := section.GetKey("password")
	if err != nil {
		return nil, errors.New("mysql config parameter(password) not found in ini section")
	}

	configKeyName, err := section.GetKey("name")
	if err != nil {
		return nil, errors.New("mysql config parameter(name) not found in ini section")
	}

	return map[string]any{
		"host":     configKeyHost.String(),
		"port":     configKeyPortInt,
		"username": configKeyUsername.String(),
		"password": configKeyPassword.String(),
		"name":     configKeyName.String(),
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
		d, err := GetDbByConfig(config)
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

// GetDbByConfig 按配置文件创建数据库实例
func GetDbByConfig(config map[string]any) (*Driver, error) {
	d := new(Driver)
	d.SetConfig(config)
	err := d.Init()
	if err != nil {
		return nil, err
	}

	return d, nil
}
