package utils

type StandardResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (response StandardResponse) Success(data interface{}) StandardResponse {
	r := StandardResponse{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	return r
}

func (response StandardResponse) Fail(code int, message string, data interface{}) StandardResponse {
	r := StandardResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return r
}
