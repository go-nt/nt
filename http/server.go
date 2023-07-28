package http

import (
	"net/http"
	"strconv"
)

type serverConfig struct {
	// 端口号
	Port uint16
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
		Port: 9999,
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
			server.config.Port = value.(uint16)
		}
	}
}

func (server *Server) AddRHandler(appName string, handler Handler) {
	if server.handlers == nil {
		server.handlers = make(map[string]Handler)
	}

	server.handlers[appName] = handler
}

// Start 启动服务
func (server *Server) Start() {

	if server.config == nil {
		server.initConfig()
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		if server.handlers != nil {

			path := r.URL.Path
			appName := ""
			i := 1
			l := len(path)
			for i < l {
				if path[i] == '/' {
					appName = path[1:i]
					break
				}
				i++
			}

			if appName != "" {
				if handler, ok := server.handlers[appName]; ok {

					req := new(Request)
					req.Init(r)

					res := new(Response)
					res.Init(rw)

					c := new(Context)
					c.Init(req, res)

					// 回收资源
					defer c.Gc()

					handler.OnRequest(c)
				}
			}
		}
	})

	err := http.ListenAndServe(":"+strconv.Itoa(int(server.config.Port)), nil)
	if err != nil {
		return
	}
}
