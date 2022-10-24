package utils

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok(data interface{}) *Response{
	return &Response{Code:200,Msg:"success",Data:data}
}

func Err(code int, msg string) *Response{
	return &Response{Code:code,Msg:msg}
}