package websocket

import (
	"errors"
	ntHttp "github.com/go-nt/nt/http"
	"net/http"
	"strconv"
)

type Handler struct {
	HandlerInterface
}

// ---------------------------------------------------------------------------------------------------------------------
// 以下方法合肥为握手阶段 HTTP 请求

// Check 检查：是否跨域，身份验证
func (handler *Handler) Check(httpCtx *ntHttp.Context) (bool, string) {
	// 校验 header 头
	// origin := httpCtx.Request.Header("Origin", "")

	// 校验 get 参数
	// token := httpCtx.Request.get("token", "")

	return true, ""
}

// OnRequest 请求
func (handler *Handler) OnRequest(httpCtx *ntHttp.Context) error {
	httpCtx.Response.Write("<p style=\"text-align:center\"><<h1>Websocket works</h1><hr>powered by <a href=\"https://www.go-nt.com\">go-nt</a></p>")
}

// OnHttpError 抿手阶段出错 HTTP 输出
func (handler *Handler) OnHttpError(httpCtx *ntHttp.Context, code int, message string) {
	httpCtx.Response.Header("Sec-Websocket-Version", "13")
	http.Error(httpCtx.Response.ResponseWriter, http.StatusText(code), code)
}

// 以上方法合肥为握手阶段 HTTP 请求
// =====================================================================================================================

// ---------------------------------------------------------------------------------------------------------------------
// 以下方法合数据传输过程中

// OnTextMessage 1: 文本消息
func (handler *Handler) OnTextMessage(wsCtx *Context, message string) error {
	return nil
}

// OnBinaryMessage 2: 二进制消息
func (handler *Handler) OnBinaryMessage(wsCtx *Context, message string) error {
	return nil
}

// OnClose close
func (handler *Handler) OnClose(wsCtx *Context) error {
	return nil
}

// OnPing ping
func (handler *Handler) OnPing(wsCtx *Context) error {
	return wsCtx.Response.Pong("pong")
}

// OnPong pong
func (handler *Handler) OnPong(wsCtx *Context) error {
	return nil
}

// OnWebsocketError 传数数据阶段出错
func (handler *Handler) OnWebsocketError(wsCtx *Context, code int, message string) {
}

// 以上方法合数据传输过程中
// =====================================================================================================================

// Start 启动
func (handler *Handler) Process(wsCtx *Context) error {

	for {
		messageType, message, err := wsCtx.Request.readMessage()
		if err != nil {
			break
		}

		switch messageType {
		case 1:
			err = handler.OnTextMessage(wsCtx, string(message))
		case 2:
			err = handler.OnBinaryMessage(wsCtx, string(message))

		case 8:
			// 关闭连接，退出处理
			_ = handler.OnClose(wsCtx)
			err = errors.New("close connection")
		case 9:
			err = handler.OnPing(wsCtx)
		case 10:
			err = handler.OnPong(wsCtx)
		default:
			err = errors.New("unsupported message type: " + strconv.Itoa(messageType))
		}

		if err != nil {
			break
		}
	}

	return nil
}
