package redis

import (
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Driver struct { // 参数配置
	config *Config
	client *redis.Client
}

// SetConfig 参数配置
func (d *Driver) SetConfig(config *Config) {
	d.config = config
}

// GetConfig 参数配置
func (d *Driver) GetConfig() *Config {
	return d.config
}

// GetClient 获取连接
func (d *Driver) GetClient() *redis.Client {
	return d.client
}

// Init 初始化
func (d *Driver) Init() error {
	d.client = redis.NewClient(&redis.Options{
		Addr:     d.config.host + ":" + strconv.Itoa(d.config.port),
		Password: d.config.password,
		DB:       d.config.db,
	})

	return nil
}
