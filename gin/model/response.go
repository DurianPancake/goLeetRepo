package model

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

// Success struct
func Success(msg string, data interface{}) *Response {
	return &Response{
		Code:    10000,
		Message: msg,
		Data:    data,
	}
}
