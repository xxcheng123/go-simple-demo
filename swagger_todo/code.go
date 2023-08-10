package main

const (
	CodeSuccess         ResponseCode = 1000
	CodeErrorUserExists ResponseCode = 1001

	CodeErrorParamError = 2001
)

var codeMap = map[ResponseCode]ResponseMessage{
	CodeSuccess:         "success",
	CodeErrorUserExists: "用户已存在",
	CodeErrorParamError: "参数不正确",
}

type ResponseCode uint16
type ResponseMessage string
type ResponseData any

type ResponseWrapper struct {
	Code    ResponseCode    `json:"code"'`          //返回状态码
	Message ResponseMessage `json:"msg"`            //提示信息
	Data    ResponseData    `json:"data,omitempty"` // 数据
}

func (rc ResponseCode) GetMsg() ResponseMessage {
	if msg, ok := codeMap[rc]; ok {
		return msg
	}
	return ""

}
