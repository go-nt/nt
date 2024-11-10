package request

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
)

type Driver struct {
	Request *http.Request
	dGet    url.Values
	dPost   url.Values
}

func (d *Driver) Init(request *http.Request) {
	d.Request = request
	d.dGet = request.URL.Query()

	err := request.ParseForm()
	if err != nil {
		return
	}

	d.dPost = request.PostForm
}

// Get 获取 string 类型的 GET 数据
func (d *Driver) Get(name string, defaultValue string) string {
	if values, ok := d.dGet[name]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}

	return defaultValue
}

// GetArray 获取 string 数组 类型的 GET 数据
func (d *Driver) GetArray(name string) []string {
	if values, ok := d.dGet[name]; ok {
		return values
	}

	return []string{}
}

// GetFormat 获取 GET 格式化数据
func (d *Driver) GetFormat(name string) *Format {
	if values, ok := d.dGet[name]; ok {
		if len(values) > 0 {
			return &Format{
				Value: values[0],
			}
		}
	}

	return &Format{}
}

// GetMap 获取 所有 GET 数据
func (d *Driver) GetMap() map[string][]string {
	return d.dGet
}

// Post 获取 string 类型的 POST 数据
func (d *Driver) Post(name string, defaultValue string) string {
	if values, ok := d.dPost[name]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}

	return defaultValue
}

// PostArray 获取 string 数组 类型的 POST 数据
func (d *Driver) PostArray(name string) []string {
	if values, ok := d.dPost[name]; ok {
		return values
	}

	return []string{}
}

// PostFormat 获取 POST 格式化数据
func (d *Driver) PostFormat(name string) *Format {
	if values, ok := d.dPost[name]; ok {
		if len(values) > 0 {
			return &Format{
				Value: values[0],
			}
		}
	}

	return &Format{}
}

// PostMap 获取 所有 POST 数据
func (d *Driver) PostMap() map[string][]string {
	return d.dPost
}

// Body 获取请求休
func (d *Driver) Body(defaultValue string) string {
	bodyBytes, err := io.ReadAll(d.Request.Body)
	if err != nil {
		return defaultValue
	}

	return string(bodyBytes)
}

// Body 获取请求休
func (d *Driver) BodyBytes(defaultValue []byte) []byte {
	bodyBytes, err := io.ReadAll(d.Request.Body)
	if err != nil {
		return defaultValue
	}

	return bodyBytes
}

// Json 获取请求休并尝试转为 JSON 格式
func (d *Driver) Json(defaultValue any) any {
	bodyBytes, err := io.ReadAll(d.Request.Body)
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
func (d *Driver) Header(name string, defaultValue string) string {
	if values, ok := d.Request.Header[name]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}

	return defaultValue
}

// HeaderFormat 获取头信息
func (d *Driver) HeaderFormat(name string) *Format {
	if values, ok := d.Request.Header[name]; ok {
		if len(values) > 0 {
			return &Format{
				Value: values[0],
			}
		}
	}

	return &Format{}
}

// HeaderArray 获取 string 数组 类型的 头信息
func (d *Driver) HeaderArray(name string) []string {
	if values, ok := d.Request.Header[name]; ok {
		return values
	}

	return []string{}
}

// HeaderMap 获取 所有 头信息
func (d *Driver) HeaderMap() map[string][]string {
	return d.Request.Header
}

// Cookie 获取 cookie 值
func (d *Driver) Cookie(name string, defaultValue string) string {
	ck, err := d.Request.Cookie(name)
	if err != nil {
		return defaultValue
	}
	return ck.Value
}

// Url 网址
func (d *Driver) Url() string {
	return d.Scheme() + "://" + d.Host() + d.Request.RequestURI
}

// RootUrl 根网址
func (d *Driver) RootUrl() string {
	return d.Scheme() + "://" + d.Host()
}

// Scheme 请求协议 "http"|"https".
func (d *Driver) Scheme() string {
	if scheme := d.Header("X-Forwarded-Proto", ""); scheme != "" {
		return scheme
	}
	if d.Request.URL.Scheme != "" {
		return d.Request.URL.Scheme
	}
	if d.Request.TLS == nil {
		return "http"
	}
	return "https"
}

// Domain 请求域名，不含端口号
func (d *Driver) Domain() string {
	if d.Request.Host != "" {
		if domain, _, err := net.SplitHostPort(d.Request.Host); err == nil {
			return domain
		}
	}
	return "localhost"
}

// Host 请求主机名，可能包含端口号
func (d *Driver) Host() string {
	if d.Request.Host != "" {
		return d.Request.Host
	}
	return "localhost"
}

// Path 请求路径
func (d *Driver) Path() string {
	return d.Request.URL.Path
}

// Method 请求方法
func (d *Driver) Method() string {
	return d.Request.Method
}

// Is 是否为指定参数 method 请求
func (d *Driver) Is(method string) bool {
	return d.Method() == method
}

// IsGet 是否为 GET 请求
func (d *Driver) IsGet() bool {
	return d.Is("GET")
}

// IsPost 是否为 POST 请求
func (d *Driver) IsPost() bool {
	return d.Is("POST")
}

// IsHead 是否为 HEAD 请求
func (d *Driver) IsHead() bool {
	return d.Is("HEAD")
}

// IsOptions 是否为 OPTIONS 请求
func (d *Driver) IsOptions() bool {
	return d.Is("OPTIONS")
}

// IsPut 是否为 PUT 请求
func (d *Driver) IsPut() bool {
	return d.Is("PUT")
}

// IsDelete 是否为 DELETE 请求
func (d *Driver) IsDelete() bool {
	return d.Is("DELETE")
}

// IsPatch 是否为 PATCH 请求
func (d *Driver) IsPatch() bool {
	return d.Is("PATCH")
}

// IsAjax 是否为 ajax 请求
func (d *Driver) IsAjax() bool {
	return d.Header("X-Requested-With", "") == "XMLHttpRequest"
}

func (d *Driver) BindGet(obj *any) error {
	return d.Bind(obj, "get")
}

func (d *Driver) BindPost(obj *any) error {
	return d.Bind(obj, "post")
}

func (d *Driver) BindBodyJson(obj *any) error {
	return d.Bind(obj, "body-json")
}

func (d *Driver) BindBodyProtobuf(obj *any) error {
	return d.Bind(obj, "body-protobuf")
}

func (d *Driver) Bind(obj *any, dsType string) error {

	switch dsType {
	case "get":
		// TODO

	case "post":
		// TODO

	case "body-json":
		bodyBytes, err := io.ReadAll(d.Request.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(bodyBytes, obj)
	case "body-protobuf":
		bodyBytes, err := io.ReadAll(d.Request.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(bodyBytes, obj)
	}
	return errors.New("unknown data source type of request")
}
