package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-ini/ini"
)

type Config struct {
	// 端口号
	port int

	// 默认处理器名称
	defaultHandlerName string
}

type Server struct {
	// 参数配置
	config *Config

	// 处理器
	handlers map[string]Handler
}

// initConfig 初始化配置
func (server *Server) initConfig() {
	server.config = &Config{
		// 端口号
		port: 9999,

		defaultHandlerName: "",
	}
}

// Config 参数配置
func (server *Server) SetConfig(config map[string]any) error {
	if server.config == nil {
		server.initConfig()
	}

	for key, value := range config {
		switch key {
		case "port":
			switch t := value.(type) {
			case int:
				if t > 0 && t < 65535 {
					server.config.port = t
				} else {
					return errors.New("http server config parameter(port) is not a valid value")
				}
			}
		case "defaultHandlerName", "default_handler_name":
			switch t := value.(type) {
			case string:
				if t != "" {
					server.config.defaultHandlerName = t
				} else {
					return errors.New("http server config parameter(default_handler_name) is not a valid value")
				}
			}
		}
	}

	return nil
}

// SetIniConfig ini 参数配置
func (server *Server) SetIniConfig(section *ini.Section) error {
	if server.config == nil {
		server.initConfig()
	}

	configKeyPort, err := section.GetKey("port")
	if err == nil {
		t, err := configKeyPort.Int()
		if err == nil && t > 0 && t < 65535 {
			server.config.port = t
		} else {
			return errors.New("http server config parameter(port) is not a valid value")
		}
	}

	configKeyDefaultHandlerName, err := section.GetKey("defaultHandlerName")
	if err == nil {
		t := configKeyDefaultHandlerName.String()
		if t != "" {
			server.config.defaultHandlerName = t
		} else {
			return errors.New("http server config parameter(defaultHandlerName) is not a valid value")
		}
	} else {
		configKeyDefaultHandlerName, err := section.GetKey("default_handler_name")
		if err == nil {
			t := configKeyDefaultHandlerName.String()
			if t != "" {
				server.config.defaultHandlerName = t
			} else {
				return errors.New("http server config parameter(default_handler_name) is not a valid value")
			}
		}
	}

	return nil
}

func (server *Server) AddHandler(handlerName string, handler Handler) {
	if server.handlers == nil {
		server.handlers = make(map[string]Handler)
	}

	server.handlers[handlerName] = handler
}

// Start 启动服务
func (server *Server) Start() {

	if server.config == nil {
		server.initConfig()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if server.handlers != nil {

			path := r.URL.Path
			handlerName := ""
			i := 1
			l := len(path)
			for i < l {
				if path[i] == '/' {
					handlerName = path[1:i]
					break
				}
				i++
			}

			if handlerName == "" && server.config.defaultHandlerName != "" {
				handlerName = server.config.defaultHandlerName
			}

			if handlerName == "" {
				_, err := w.Write([]byte("<a href=\"https://www.go-nt.com\" target=\"_blank\">GO-NT</a> framework!"))
				if err != nil {
					return
				}
			} else {
				if handler, ok := server.handlers[handlerName]; ok {

					c := new(Context)
					c.Init(r, w)

					handler.OnRequest(c)
				} else {
					_, err := w.Write([]byte("Handler(" + handlerName + ") does not exist"))
					if err != nil {
						return
					}
				}
			}
		}
	})

	err := http.ListenAndServe(":"+strconv.Itoa(int(server.config.port)), nil)
	if err != nil {
		return
	}
}
