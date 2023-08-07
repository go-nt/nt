package websocket

import "net"

type Context struct {
	Id       uint64
	Conn     net.Conn
	Request  *Request
	Response *Response
	Server   *Server
}

// Init 初始化
func (c *Context) Init(id uint64, conn net.Conn, server *Server) {
	c.Id = id
	c.Conn = conn
	c.Server = server

	req := new(Request)
	req.Init(conn, server)
	c.Request = req

	res := new(Response)
	res.Init(conn, server)
	c.Response = res
}

// Gc 回收资源
func (c *Context) Gc() {

	c.Server.Clients.Delete(c.Id)

	err := c.Conn.Close()
	if err != nil {
		return
	}
}
