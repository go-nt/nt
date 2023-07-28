package websocket

import (
	ntHttp "github.com/go-nt/nt/http"
)

type HandlerInterface interface {

	// ---------------------------------------------------------------------------------------------------------------------
	// 以下方法合肥为握手阶段 HTTP 请求

	// OnRequest 默认请求, 未升级到 websocket HTTP 请求
	OnRequest(*ntHttp.Context)

	// Check 检查：是否跨域，身份验证
	Check(*ntHttp.Context) (bool, string)

	// OnHttpError 出错
	OnHttpError(*ntHttp.Context, int, string)

	// 以上方法合肥为握手阶段 HTTP 请求
	// =====================================================================================================================

	// ---------------------------------------------------------------------------------------------------------------------
	// 以下方法合数据传输过程中

	// Ping 心跳包
	OnPing(*Context)

	// OnHWebsocketError 出错
	OnHWebsocketError(*Context) bool

	// OnMessage 接收到消处
	OnMessage(*Context) string

	// OnError 出错
	OnWebsocketError(*Context, int, string)

	// 以上方法合数据传输过程中
	// =====================================================================================================================

}
