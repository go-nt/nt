package http

import (
	"net/http"
	"strconv"
)

type serverConfig struct {
	// 端口号
	port uint16

	// 默认处理器名称
	defaultHandlerMame string
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

		defaultHandlerMame: "",
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
			server.config.port = value.(uint16)
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

			if handlerName == "" && server.config.defaultHandlerMame != "" {
				handlerName = server.config.defaultHandlerMame
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

					// 回收资源
					defer c.Gc()

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
