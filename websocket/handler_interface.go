package websocket

import (
	ntHttp "github.com/go-nt/nt/http"
)

type HandlerInterface interface {

	// ---------------------------------------------------------------------------------------------------------------------
	// 以下方法合肥为握手阶段 HTTP 请求

	// OnRequest 默认请求, 未升级到 websocket HTTP 请求
	OnRequest(*ntHttp.Context) error

	// Check 检查：是否跨域，身份验证
	Check(*ntHttp.Context) (bool, string)

	// OnHttpError 出错
	OnHttpError(*ntHttp.Context, int, string)

	// 以上方法合肥为握手阶段 HTTP 请求
	// =====================================================================================================================

	// ---------------------------------------------------------------------------------------------------------------------
	// 以下方法合数据传输过程中

	// OnTextMessage 1: 文本消息
	OnTextMessage(*Context, string) error

	// OnBinaryMessage 2: 二进制消息
	OnBinaryMessage(*Context, string) error

	// OnClose close
	OnClose(*Context) error

	// OnPing ping
	OnPing(*Context) error

	// OnPong pong
	OnPong(*Context) error

	// OnWebsocketError 出错
	OnWebsocketError(*Context, int, string)

	// 以上方法合数据传输过程中
	// =====================================================================================================================

	// Start 启动
	Process(wsCtx *Context) error
}
