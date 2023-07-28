package websocket

import "net"

type Context struct {
	Conn     net.Conn
	Request  *Request
	Response *Response
}

// Init 初始化
func (c *Context) Init(conn net.Conn) {
	c.Conn = conn

	req := new(Request)
	req.Init(conn)
	c.Request = req

	res := new(Response)
	res.Init(conn)
	c.Response = res
}

// Gc 回收资源
func (c *Context) Gc() {

}
