package helper

type BaseHttpResponse struct {
	Result     any    `json:"result"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`

	Error any `json:"error"`
}

func GenerateBaseResponse(success bool, message string, statusCode int, result any) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		StatusCode: statusCode,
		Message:    message,
	}
}

func GenerateBaseResponseWithError(result any, statusCode int, err error, message string) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    false,
		StatusCode: statusCode,
		Error:      err.Error(),
		Message:    message,
	}
}
