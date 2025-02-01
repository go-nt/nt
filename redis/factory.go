package redis

import (
	"errors"
)

var drivers map[string]*Driver

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
