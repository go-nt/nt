package websocket

import (
	"net"
	"strconv"
)

type Request struct {
	Conn net.Conn
}

// Init 初始化
func (r *Request) Init(Conn net.Conn) {
	r.Conn = Conn
}

// RemoteAddr 请求者 地址
func (r *Request) RemoteAddr() string {
	return r.Conn.RemoteAddr().String()
}

// RemoteIp 请求者 IP
func (r *Request) RemoteIp() string {
	addr := r.Conn.RemoteAddr().String()

	ip := addr
	i := 7
	l := len(addr)
	for i < l {
		if addr[i] == ':' {
			ip = addr[0:i]
			break
		}
		i++
	}

	return ip
}

// RemotePort 请求者 端口号
func (r *Request) RemotePort() uint16 {
	addr := r.Conn.RemoteAddr().String()

	var port uint16 = 0
	i := 7
	l := len(addr)
	for i < l {
		if addr[i] == ':' {
			val, err := strconv.ParseInt(addr[i:l], 10, 16)
			if err == nil {
				port = uint16(val)
			}
			break
		}
		i++
	}

	return port
}
