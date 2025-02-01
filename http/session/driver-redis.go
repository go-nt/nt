package session

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	ntHttp "github.com/go-nt/nt/http"
	"github.com/go-nt/nt/redis"
	"github.com/google/uuid"
)

type DriverRedis struct {
	config *Config
	ctx    *ntHttp.Context
	id     string
	data   map[string]any
	redis  *redis.Driver
}

// Init 初始化
func (d *DriverRedis) Init(config *Config, ctx *ntHttp.Context) {
	d.config = config
	d.ctx = ctx

	redis, err := redis.GetRedis(d.config.redis)
	if err != nil {
		fmt.Println("session redis driver error: " + err.Error())
		return
	}

	d.redis = redis

	d.id = ctx.Request.Cookie(config.name, "")
	if d.id != "" {
		err := uuid.Validate(d.id)
		if err == nil {
			dataJson, err := d.redis.GetClient().Get(context.TODO(), "session:"+d.id).Bytes()
			if err == nil {
				var data map[string]any
				err = json.Unmarshal(dataJson, &data)
				if err == nil {
					d.data = data
				}
			}
		} else {
			d.id = ""
		}
	}

	if d.id == "" {
		id := uuid.New()
		d.id = id.String()
	}

	// 将 session  id 写入 cookie
	cookie := &http.Cookie{
		Name:     config.name,
		Value:    d.id,
		Expires:  time.Now().Add(time.Duration(config.expire) * time.Second),
		HttpOnly: true,
	}

	ctx.Response.Cookie(cookie)
}

// 数据持久化
func (d *DriverRedis) Save() {
	if d.id != "" {
		redisKey := "session:" + d.id
		if len(d.data) > 0 {
			data, _ := json.Marshal(d.data)
			d.redis.GetClient().Set(context.TODO(), redisKey, data, time.Duration(d.config.expire))
		} else {
			d.redis.GetClient().Del(context.TODO(), redisKey)
		}
	}
}

// 获取ID
func (d *DriverRedis) GetId() string {
	return d.id
}

// Get 获取 session 值
func (d *DriverRedis) Get(name string) any {
	value, _ := d.data[name]
	return value
}

// Set 向 session 中写入
func (d *DriverRedis) Set(name string, value any) {
	d.data[name] = value
}

// Has 是否已设置指定名称的 session
func (d *DriverRedis) Has(name string) bool {
	_, exists := d.data[name]
	return exists
}

// Delete 删除指定锓名的 session
func (d *DriverRedis) Delete(name string) any {
	val, exists := d.data[name]
	if exists {
		delete(d.data, name)
	}

	return val
}

// Wipe 清空 session
func (d *DriverRedis) Wipe() {
	d.data = nil
}
