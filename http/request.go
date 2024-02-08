package http

import (
	"encoding/json"
	"github.com/go-nt/nt/http/request"
	"io"
	"net"
	"net/http"
	"net/url"
)

type Request struct {
	Request *http.Request
	dGet    url.Values
	dPost   url.Values
}

func (r *Request) Init(request *http.Request) {
	r.Request = request
	r.dGet = request.URL.Query()

	err := request.ParseForm()
	if err != nil {
		return
	}

	r.dPost = request.PostForm
}

// Get 获取 string 类型的 GET 数据
func (r *Request) Get(name string, defaultValue string) string {
	if values, ok := r.dGet[name]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}

	return defaultValue
}

// GetArray 获取 string 数组 类型的 GET 数据
func (r *Request) GetArray(name string) []string {
	if values, ok := r.dGet[name]; ok {
		return values
	}

	return []string{}
}

// GetFormat 获取 GET 格式化数据
func (r *Request) GetFormat(name string) *request.Format {
	if values, ok := r.dGet[name]; ok {
		if len(values) > 0 {
			return &request.Format{
				Value: values[0],
			}
		}
	}

	return &request.Format{}
}

// GetMap 获取 所有 GET 数据
func (r *Request) GetMap() map[string][]string {
	return r.dGet
}

// Post 获取 string 类型的 POST 数据
func (r *Request) Post(name string, defaultValue string) string {
	if values, ok := r.dPost[name]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}

	return defaultValue
}

// PostArray 获取 string 数组 类型的 POST 数据
func (r *Request) PostArray(name string) []string {
	if values, ok := r.dPost[name]; ok {
		return values
	}

	return []string{}
}

// PostFormat 获取 POST 格式化数据
func (r *Request) PostFormat(name string) *request.Format {
	if values, ok := r.dPost[name]; ok {
		if len(values) > 0 {
			return &request.Format{
				Value: values[0],
			}
		}
	}

	return &request.Format{}
}

// PostMap 获取 所有 POST 数据
func (r *Request) PostMap() map[string][]string {
	return r.dPost
}

// Body 获取请求休
func (r *Request) Body(defaultValue string) string {
	bodyBytes, err := io.ReadAll(r.Request.Body)
	if err != nil {
		return defaultValue
	}

	return string(bodyBytes)
}

// Json 获取请求休并尝试转为 JSON 格式
func (r *Request) Json(defaultValue any) any {
	bodyBytes, err := io.ReadAll(r.Request.Body)
	if err != nil {
		return defaultValue
	}

	var data any
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return defaultValue
	}

	return data
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

// HeaderFormat 获取头信息
func (r *Request) HeaderFormat(name string) *request.Format {
	if values, ok := r.Request.Header[name]; ok {
		if len(values) > 0 {
			return &request.Format{
				Value: values[0],
			}
		}
	}

	return &request.Format{}
}

// HeaderArray 获取 string 数组 类型的 头信息
func (r *Request) HeaderArray(name string) []string {
	if values, ok := r.Request.Header[name]; ok {
		return values
	}

	return []string{}
}

// HeaderMap 获取 所有 头信息
func (r *Request) HeaderMap() map[string][]string {
	return r.Request.Header
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
