package session

import (
	ntHttp "github.com/go-nt/nt/http"
)

type Driver interface {

	// Init 初始化
	Init(config *Config, ctx *ntHttp.Context)

	// 获取 session ID
	GetId() string

	Save()

	// Get 获取 session 值
	Get(key string) any

	// Set 向 session 中写入
	Set(key string, val any)

	// Has 是否已设置指定名称的 session
	Has(key string) bool

	// Delete 删除指定锓名的 session
	Delete(key string) any

	// Wipe 清空 session
	Wipe()
}
