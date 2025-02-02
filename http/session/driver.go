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
	Get(name string) any

	// Get 获取 session 值
	GetFormat(name string) *Format

	// Set 向 session 中写入
	Set(name string, value any)

	// Has 是否已设置指定名称的 session
	Has(name string) bool

	// Delete 删除指定锓名的 session
	Delete(name string) any

	// Wipe 清空 session
	Wipe()
}
