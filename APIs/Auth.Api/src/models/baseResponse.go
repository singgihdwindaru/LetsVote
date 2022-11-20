package models

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func HttpResponse(code int, message string, data interface{}) WebResponse {
	response := WebResponse{
		Code:    code,
		Message: message,
		Payload: data,
	}
	return response
}
