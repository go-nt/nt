package http

import (
	"net/http"
)

type Context struct {
	Request  *Request
	Response *Response
}

// Init 初始化
func (c *Context) Init(r *http.Request, w http.ResponseWriter) {
	req := new(Request)
	req.Init(r)

	res := new(Response)
	res.Init(w)

	c.Request = req
	c.Response = res
}

func (c *Context) GetDb() {

}

func (c *Context) GetRedis() {

}

func (c *Context) GetSession() {

}

// Gc 回收资源
func (c *Context) Gc() {

}
