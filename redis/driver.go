package redis

import (
	"strconv"

	"github.com/go-redis/redis/v8"
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

type Driver struct { // 参数配置
	config *Config
	client *redis.Client
}

// initConfig 初始化配置
func (d *Driver) initConfig() {
	d.config = &Config{
		host:     "127.0.0.1",
		port:     6379,
		password: "",
		db:       0,
	}
}

// SetConfig 参数配置
func (d *Driver) SetConfig(config map[string]any) {
	if d.config == nil {
		d.initConfig()
	}

	for key, value := range config {
		switch key {
		case "host":
			switch t := value.(type) {
			case string:
				d.config.host = t
			}
		case "port":
			switch t := value.(type) {
			case int:
				if t > 0 && t < 65535 {
					d.config.port = t
				}
			}
		case "password":
			switch t := value.(type) {
			case string:
				d.config.password = t
			}
		case "db":
			switch t := value.(type) {
			case int:
				d.config.db = t
			}
		}
	}
}

// GetConfig 参数配置
func (d *Driver) GetConfig() *Config {
	if d.config == nil {
		d.initConfig()
	}

	return d.config
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

// GetClient 获取连接
func (d *Driver) GetClient() *redis.Client {
	return d.client
}
