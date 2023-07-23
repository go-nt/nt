package http

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
)

type Request struct {
	Request    *http.Request
	dataOfGet  url.Values
	dataOfPost url.Values
}

func (r *Request) Init(request *http.Request) {
	r.Request = request
	r.dataOfGet = request.URL.Query()

	err := request.ParseForm()
	if err != nil {
		return
	}

	r.dataOfPost = request.PostForm
}

// GetAll 获取 所有 GET 数据
func (r *Request) GetAll() map[string][]string {
	return r.dataOfGet
}

// Get 获取 string 类型的 GET 数据
func (r *Request) Get(name string, defaultValue string) string {
	if values, ok := r.dataOfGet[name]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}

	return defaultValue
}

// GetArray 获取 string 类型的 GET 数据数组
func (r *Request) GetArray(name string) []string {
	if values, ok := r.dataOfGet[name]; ok {
		return values
	}

	return []string{}
}

// GetByte 获取 byte 类型的 GET 数据
func (r *Request) GetByte(name string, defaultValue byte) byte {
	v := r.Get(name, "")
	if len(v) != 1 {
		return defaultValue
	}

	return v[0]
}

// GetInt 获取 int 类型的 GET 数据
func (r *Request) GetInt(name string, defaultValue int) int {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	intVal, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}

	return intVal
}

// GetInt8 获取 int8 类型的 GET 数据
func (r *Request) GetInt8(name string, defaultValue int8) int8 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return defaultValue
	}

	return int8(val)
}

// GetInt16 获取 int16 类型的 GET 数据
func (r *Request) GetInt16(name string, defaultValue int16) int16 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return defaultValue
	}

	return int16(val)
}

// GetInt32 获取 int32 类型的 GET 数据
func (r *Request) GetInt32(name string, defaultValue int32) int32 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return defaultValue
	}

	return int32(val)
}

// GetInt64 获取 int64 类型的 GET 数据
func (r *Request) GetInt64(name string, defaultValue int64) int64 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultValue
	}

	return val
}

// GetUint 获取 uint 类型的 GET 数据
func (r *Request) GetUint(name string, defaultValue uint) uint {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return defaultValue
	}

	return uint(val)
}

// GetUint8 获取 uint8 类型的 GET 数据
func (r *Request) GetUint8(name string, defaultValue uint8) uint8 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return defaultValue
	}

	return uint8(val)
}

// GetUint16 获取 uint16 类型的 GET 数据
func (r *Request) GetUint16(name string, defaultValue uint16) uint16 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		return defaultValue
	}

	return uint16(val)
}

// GetUint32 获取 uint32 类型的 GET 数据
func (r *Request) GetUint32(name string, defaultValue uint32) uint32 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return defaultValue
	}

	return uint32(val)
}

// GetUnt64 获取 uint64 类型的 GET 数据
func (r *Request) GetUnt64(name string, defaultValue uint64) uint64 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return defaultValue
	}

	return val
}

// GetFloat32 获取 float32 类型的 GET 数据
func (r *Request) GetFloat32(name string, defaultValue float32) float32 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return defaultValue
	}

	return float32(val)
}

// GetFloat64 获取 float64 类型的 GET 数据
func (r *Request) GetFloat64(name string, defaultValue float64) float64 {
	v := r.Get(name, "")
	if v == "" {
		return defaultValue
	}

	val, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return defaultValue
	}

	return val
}

// Post 获取 string 类型的 POST 数据
func (r *Request) Post(name string, defaultValue string) string {
	if values, ok := r.dataOfPost[name]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}

	return defaultValue
}

// Header 获取头信息
func (r *Request) Header(name string, defaultValue string) string {
	if values, ok := r.Request.Header[name]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}

	return defaultValue
}

// Cookie 获取 cookie 值
func (r *Request) Cookie(name string, defaultValue string) string {
	ck, err := r.Request.Cookie(name)
	if err != nil {
		return defaultValue
	}
	return ck.Value
}

// Url 网址
func (r *Request) Url() string {
	return r.Scheme() + "://" + r.Host() + r.Request.RequestURI
}

// RootUrl 根网址
func (r *Request) RootUrl() string {
	return r.Scheme() + "://" + r.Host()
}

// Scheme 请求协议 "http"|"https".
func (r *Request) Scheme() string {
	if scheme := r.Header("X-Forwarded-Proto", ""); scheme != "" {
		return scheme
	}
	if r.Request.URL.Scheme != "" {
		return r.Request.URL.Scheme
	}
	if r.Request.TLS == nil {
		return "http"
	}
	return "https"
}

// Domain 请求域名，不含端口号
func (r *Request) Domain() string {
	if r.Request.Host != "" {
		if domain, _, err := net.SplitHostPort(r.Request.Host); err == nil {
			return domain
		}
	}
	return "localhost"
}

// Host 请求主机名，可能包含端口号
func (r *Request) Host() string {
	if r.Request.Host != "" {
		return r.Request.Host
	}
	return "localhost"
}

// Path 请求路径
func (r *Request) Path() string {
	return r.Request.URL.Path
}

// Method 请求方法
func (r *Request) Method() string {
	return r.Request.Method
}

// Is 是否为指定参数 method 请求
func (r *Request) Is(method string) bool {
	return r.Method() == method
}

// IsGet 是否为 GET 请求
func (r *Request) IsGet() bool {
	return r.Is("GET")
}

// IsPost 是否为 POST 请求
func (r *Request) IsPost() bool {
	return r.Is("POST")
}

// IsHead 是否为 HEAD 请求
func (r *Request) IsHead() bool {
	return r.Is("HEAD")
}

// IsOptions 是否为 OPTIONS 请求
func (r *Request) IsOptions() bool {
	return r.Is("OPTIONS")
}

// IsPut 是否为 PUT 请求
func (r *Request) IsPut() bool {
	return r.Is("PUT")
}

// IsDelete 是否为 DELETE 请求
func (r *Request) IsDelete() bool {
	return r.Is("DELETE")
}

// IsPatch 是否为 PATCH 请求
func (r *Request) IsPatch() bool {
	return r.Is("PATCH")
}

// IsAjax 是否为 ajax 请求
func (r *Request) IsAjax() bool {
	return r.Header("X-Requested-With", "") == "XMLHttpRequest"
}
