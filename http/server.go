package http

import (
	"net/http"
)

type Server struct {
	handlers map[string]Handler
}

// Start 启动服务
func (server *Server) Start(port int) {

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

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		return
	}
}

func (server *Server) AddRHandler(appName string, handler Handler) *Server {
	if server.handlers == nil {
		server.handlers = make(map[string]Handler)
	}

	server.handlers[appName] = handler
	return server
}
