package response

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Driver struct {
	http.ResponseWriter
	data map[string]any
}

// Init 初始化
func (d *Driver) Init(rw http.ResponseWriter) {
	d.ResponseWriter = rw
	d.data = make(map[string]any)
}

// Header 输出头你息
func (d *Driver) Header(name string, value string) {
	d.ResponseWriter.Header().Set(name, value)
}

// Write 输出内容
func (d *Driver) Write(content string) {
	_, _ = d.ResponseWriter.Write([]byte(content))
}

// Set 设置数据
func (d *Driver) Set(name string, value any) {
	d.data[name] = value
}

// Json 输出JSON
func (d *Driver) Json() {
	content, _ := json.Marshal(d.data)
	_, _ = d.ResponseWriter.Write([]byte(content))
}

// Display 显示模板
func (d *Driver) Display(filenames ...string) {
	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		return
	}

	_ = tmpl.Execute(d.ResponseWriter, d.data)
}
