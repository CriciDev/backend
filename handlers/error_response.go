package handlers

type ErrorResponse struct {
	Code    int
	Message string
}

func NewError(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}
