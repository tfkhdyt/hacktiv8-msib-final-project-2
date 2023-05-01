package errs

import "net/http"

type MessageErr interface {
	Message() string
	StatusCode() int
	Error() string
}

type MessageErrData struct {
	ErrMessage    string `json:"message" example:"This is an error message"`
	ErrStatusCode int    `json:"statusCode" example:"400"`
	ErrError      string `json:"error" example:"BAD_REQUEST"`
}

func (e *MessageErrData) Message() string {
	return e.ErrMessage
}

func (e *MessageErrData) StatusCode() int {
	return e.ErrStatusCode
}

func (e *MessageErrData) Error() string {
	return e.ErrError
}

func NewInternalServerError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessableEntity(message string) MessageErr {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusUnprocessableEntity,
		ErrError:      "INVALID_REQUEST_BODY",
	}
}

func NewBadRequest(message string) MessageErr {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusBadRequest,
		ErrError:      "BAD_REQUEST",
	}
}

func NewNotFound(message string) MessageErr {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusNotFound,
		ErrError:      "DATA_NOT_FOUND",
	}
}

func NewUnauthenticated(message string) MessageErr {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusUnauthorized,
		ErrError:      "UNAUTHENTICATED",
	}
}

func NewUnauthorized(message string) MessageErr {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusForbidden,
		ErrError:      "UNAUTHORIZED",
	}
}
