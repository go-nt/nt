package websocket

import (
	"net"
)

type Response struct {
	Conn net.Conn
}

// Init 初始化
func (r *Response) Init(Conn net.Conn) {
	r.Conn = Conn
}

// Ping 发送 ping
func (r *Response) Ping(content string) {

}

// Pong 响应 ping
func (r *Response) Pong(content string) {

}

// WriteMessage 写数据
func (r *Response) Write(content string) {

}

// Close 关闭链接
func (r *Response) Close(content string) {

}
