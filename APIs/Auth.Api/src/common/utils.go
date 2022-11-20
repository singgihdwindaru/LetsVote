package common

import "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"

func HttpResponse(code int, message string, data interface{}) models.WebResponse {
	response := models.WebResponse{
		Code:    code,
		Message: message,
		Payload: data,
	}
	return response
}
