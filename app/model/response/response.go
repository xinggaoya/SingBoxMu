package response

/**
  @author: XingGao
  @date: 2024/8/9
**/

// ResInfo 定义 HTTP 接口返回结构
type ResInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 成功返回
func Success(data interface{}) ResInfo {
	response := ResInfo{
		Code: 10000,
		Msg:  "success",
		Data: data,
	}
	return response
}

// Error 失败返回
func Error(msg string) ResInfo {
	response := ResInfo{
		Code: 10001,
		Msg:  msg,
	}
	return response
}
