package http

import (
	"net/http"
	"strconv"
)

type serverConfig struct {
	// 端口号
	port uint16

	// 默认处理器名称
	defaultHandlerName string
}

type Server struct {
	// 参数配置
	config *serverConfig

	// 处理器
	handlers map[string]Handler
}

// initConfig 初始化配置
func (server *Server) initConfig() {
	server.config = &serverConfig{
		// 端口号
		port: 9999,

		defaultHandlerName: "",
	}
}

// Config 参数配置
func (server *Server) Config(config map[string]any) {
	if server.config == nil {
		server.initConfig()
	}

	for key, value := range config {
		switch key {
		case "port":
			switch t := value.(type) {
			case int:
				if t > 0 && t < 65535 {
					server.config.port = uint16(t)
				}
			case uint16:
				if t > 0 {
					server.config.port = t
				}
			}
		case "defaultHandlerName", "default_handler_name":
			switch t := value.(type) {
			case string:
				server.config.defaultHandlerName = t
			}
		}
	}
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
