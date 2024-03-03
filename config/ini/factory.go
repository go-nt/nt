package ini

import (
	"errors"
)

var configs map[string]*Config

func GetConfig(name string) (*Config, error) {
	if configs == nil {
		configs = make(map[string]*Config)
	}

	config, ok := configs[name]
	if ok {
		return config, nil
	}

	return nil, errors.New("ini config (" + name + ") not found")
}

func SetConfig(name string, path string) {
	if configs == nil {
		configs = make(map[string]*Config)
	}

	config := new(Config)
	config.SetPath(path)
	configs[name] = config
}
