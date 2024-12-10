package utils

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (response Response) Success(data interface{}) Response {
	r := Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	return r
}
