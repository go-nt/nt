package session

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	ntHttp "github.com/go-nt/nt/http"
	"github.com/go-nt/nt/util/fs/dir"
	"github.com/go-nt/nt/util/fs/file"
	"github.com/google/uuid"
)

type DriverFile struct {
	config *Config
	ctx    *ntHttp.Context
	id     string
	data   map[string]any
	path   string
}

// Init 初始化
func (d *DriverFile) Init(config *Config, ctx *ntHttp.Context) {
	d.config = config
	d.ctx = ctx
	d.path = ".session"

	d.id = ctx.Request.Cookie(config.name, "")
	if d.id != "" {
		err := uuid.Validate(d.id)
		if err == nil {
			path := filepath.Join("data", d.path, d.id)
			ok, _ := file.IsFile(path)
			if ok {
				dataJson, err := os.ReadFile(path)
				if err == nil {
					var data map[string]any
					err = json.Unmarshal(dataJson, &data)
					if err == nil {
						d.data = data
					}
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
func (d *DriverFile) Save() {
	if d.id != "" {
		path := filepath.Join("data", d.path)
		ok, _ := dir.IsDir(path)
		if !ok {
			dir.Make(path)
		}

		path = filepath.Join(path, d.id)

		if len(d.data) > 0 {
			data, _ := json.Marshal(d.data)
			os.WriteFile(path, []byte(data), os.ModePerm)
		} else {
			os.Remove(path)
		}
	}
}

// 获取 sessionID
func (d *DriverFile) GetId() string {
	return d.id
}

// Get 获取 session 值
func (d *DriverFile) Get(name string) any {
	value, _ := d.data[name]
	return value
}

// GetFormat 获取 GET 格式化数据
func (d *DriverFile) GetFormat(name string) *Format {
	if value, ok := d.data[name]; ok {
		return &Format{
			Value: value,
		}
	}

	return &Format{}
}

// Set 向 session 中写入
func (d *DriverFile) Set(name string, value any) {
	d.data[name] = value
}

// Has 是否已设置指定名称的 session
func (d *DriverFile) Has(name string) bool {
	_, exists := d.data[name]
	return exists
}

// Delete 删除指定锓名的 session
func (d *DriverFile) Delete(name string) any {
	value, exists := d.data[name]
	if exists {
		delete(d.data, name)
	}

	return value
}

// Wipe 清空 session
func (d *DriverFile) Wipe() {
	d.data = nil
}
