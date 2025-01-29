package request

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
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

func (d *Driver) GetBind(ptr any) error {
	return d.bind(ptr, "get")
}

func (d *Driver) PostBind(ptr any) error {
	return d.bind(ptr, "post")
}

func (d *Driver) BodyJsonBind(ptr any) error {
	return d.bind(ptr, "body-json")
}

func (d *Driver) bind(ptr any, dsType string) error {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Ptr {
		return errors.New("request bind error: param of ptr is not a pointer")
	}

	rv = rv.Elem()

	// json 中为切片的情况
	if dsType == "body-json" {
		bodyBytes, err := io.ReadAll(d.Request.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(bodyBytes, rv)
		if err != nil {
			return err
		}

		return nil
	}

	return d.bindValue(rv, dsType)
}

// 绑定
func (d *Driver) bindValue(rv reflect.Value, dsType string) error {
	rt := rv.Type()
	if rt.Kind() != reflect.Struct {
		return errors.New("request bind error: param of ptr is not a struct")
	}

	for i := 0; i < rv.NumField(); i++ {
		rvField := rv.Field(i)
		rtField := rt.Field(i)

		tag := rtField.Tag.Get("bind")
		if tag == "-" {
			continue
		}

		if len(rtField.Name) == 0 || !rvField.CanSet() {
			continue
		}

		if tag == "" {
			tag = rtField.Name
		}

		rvFieldKind := rvField.Kind()

		switch rvFieldKind {
		case reflect.String,
			reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:

			var val string
			switch dsType {
			case "get":
				val = d.Get(tag, "")
			case "post":
				val = d.Post(tag, "")
			case "header":
				val = d.Header(tag, "")
			case "cookie":
				val = d.Cookie(tag, "")
			}

			switch rvFieldKind {
			case reflect.String:
				rvField.SetString(val)
			case reflect.Bool:
				if val == "true" || val == "1" {
					rvField.SetBool(true)
				} else {
					rvField.SetBool(false)
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valInt64, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					rvField.SetInt(valInt64)
				}
			case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				valUint64, err := strconv.ParseUint(val, 10, 64)
				if err == nil {
					rvField.SetUint(valUint64)
				}
			case reflect.Float32, reflect.Float64:
				valFloat64, err := strconv.ParseFloat(val, 64)
				if err == nil {
					rvField.SetFloat(valFloat64)
				}
			}

		case reflect.Slice:

			var vals []string
			switch dsType {
			case "get":
				vals = d.GetArray(tag + "[]")
			case "post":
				vals = d.PostArray(tag + "[]")
			case "header":
				vals = d.HeaderArray(tag + "[]")
			case "cookie":
				vals = strings.Split(d.Cookie(tag, ""), ",")
			}

			if len(vals) > 0 {
				newSlice := reflect.MakeSlice(rtField.Type, 0, len(vals))
				switch rtField.Type.Elem().Kind() {
				case reflect.String:
					newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(vals))
				case reflect.Bool:
					for _, val := range vals {
						if val == "true" || val == "1" {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(true))
						} else {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(false))
						}
					}
				case reflect.Int:
					for _, val := range vals {
						valInt64, err := strconv.ParseInt(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(int(valInt64)))
						}
					}
				case reflect.Int8:
					for _, val := range vals {
						valInt64, err := strconv.ParseInt(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(int8(valInt64)))
						}
					}
				case reflect.Int16:
					for _, val := range vals {
						valInt64, err := strconv.ParseInt(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(int16(valInt64)))
						}
					}
				case reflect.Int32:
					for _, val := range vals {
						valInt64, err := strconv.ParseInt(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(int32(valInt64)))
						}
					}
				case reflect.Int64:
					for _, val := range vals {
						valInt64, err := strconv.ParseInt(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(valInt64))
						}
					}
				case reflect.Uint:
					for _, val := range vals {
						valUint64, err := strconv.ParseUint(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(uint(valUint64)))
						}
					}
				case reflect.Uint8:
					for _, val := range vals {
						valUint64, err := strconv.ParseUint(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(uint8(valUint64)))
						}
					}
				case reflect.Uint16:
					for _, val := range vals {
						valUint64, err := strconv.ParseUint(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(uint16(valUint64)))
						}
					}
				case reflect.Uint32:
					for _, val := range vals {
						valUint64, err := strconv.ParseUint(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(uint32(valUint64)))
						}
					}
				case reflect.Uint64:
					for _, val := range vals {
						valUint64, err := strconv.ParseUint(val, 10, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(valUint64))
						}
					}
				case reflect.Float32:
					for _, val := range vals {
						valFloat64, err := strconv.ParseFloat(val, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(float32(valFloat64)))
						}
					}
				case reflect.Float64:
					for _, val := range vals {
						valFloat64, err := strconv.ParseFloat(val, 64)
						if err == nil {
							newSlice = reflect.Append(newSlice, reflect.ValueOf(valFloat64))
						}
					}
				}

				rvField.Set(newSlice)
			}

		}
	}

	return nil
}
