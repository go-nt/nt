package http

import (
	"encoding/json"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
}

type Response struct {
	ResponseWriter *ResponseWriter
	data           map[string]any
}

// Init 初始化
func (r *Response) Init(rw http.ResponseWriter) {
	r.ResponseWriter = new(ResponseWriter)
	r.ResponseWriter.ResponseWriter = rw
	r.data = make(map[string]any)
}

// Header 输出头你息
func (r *Response) Header(name string, value string) {
	r.ResponseWriter.Header().Set(name, value)
}

// Write 输出内容
func (r *Response) Write(content string) {
	_, _ = r.ResponseWriter.Write([]byte(content))
}

// Set 设置数据
func (r *Response) Set(name string, value any) {
	r.data[name] = value
}

// Json 输出JSON
func (r *Response) Json() {
	content, _ := json.Marshal(r.data)
	_, _ = r.ResponseWriter.Write([]byte(content))
}
