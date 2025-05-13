package rest

// resp.code 表示业务状态, 用于进一步细分 http.code.
// http.code 表示接口整体状态, 如 404 not found.
// resp.code 进一步区分业务状态, 如 20001 订单已退款, 表示接口正常但业务自身逻辑失败.
type (
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"msg"`
		Data    interface{} `json:"data"`
	}
)

// 请求成功的响应
func SucResponse(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}

// 请求成功的响应
func FailResponse(code int, msg string) *Response {
	return &Response{
		Code:    code,
		Message: msg,
	}
}
