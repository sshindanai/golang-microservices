package errors

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AnErr    string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.AStatus
}

func (e *apiError) Message() string {
	return e.AMessage
}

func (e *apiError) Error() string {
	return e.AnErr
}

func NewNotFoundApiError(msg string) ApiError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: msg,
	}
}

func NewApiError(statusCode int, msg string) ApiError {
	return &apiError{
		AStatus:  statusCode,
		AMessage: msg,
	}
}

func NewInternalServerError(msg string) ApiError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: msg,
	}
}

func NewBadRequestError(msg string) ApiError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: msg,
	}
}
