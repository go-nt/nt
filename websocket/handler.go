package websocket

import (
	ntHttp "github.com/go-nt/nt/http"
	"net/http"
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
func (handler *Handler) OnRequest(httpCtx *ntHttp.Context) {
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

// Ping 连接保持
func (handler *Handler) OnPing(*Context) {
	//return "pong"
}

// Ping 连接保持
func (handler *Handler) OnMessage(*Context) string {
	return "pong"
}

// OnWebsocketError 传数数据阶段出错
func (handler *Handler) OnWebsocketError(wsCtx *Context, code int, message string) {

}

// 以上方法合数据传输过程中
// =====================================================================================================================
