package ini

import (
	"errors"
	"github.com/go-ini/ini"
)

type Config struct {
	// 配置文件路径
	path string

	// 暂存配置文件
	file *ini.File
}

// SetPath 设置配置文件路径
func (c *Config) SetPath(path string) {
	c.path = path
}

// GetPath 获取配置文件路径
func (c *Config) GetPath(path string) string {
	return path
}

func (c *Config) GetFile() (*ini.File, error) {
	if c.file == nil {
		if c.path == "" {
			return nil, errors.New("config ini path does not set")
		}

		f, err := ini.Load(c.path)
		if err != nil {
			return nil, err
		}
		c.file = f
	}

	return c.file, nil
}

func (c *Config) GetSection(section string) (*ini.Section, error) {
	f, err := c.GetFile()
	if err != nil {
		return nil, err
	}
	return f.Section(section), nil
}

func (c *Config) GetKey(section string, key string) (*ini.Key, error) {
	s, err := c.GetSection(section)
	if err != nil {
		return nil, err
	}
	return s.Key(key), nil
}
