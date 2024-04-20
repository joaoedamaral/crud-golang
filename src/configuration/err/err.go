package err

import "net/http"

type Err struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewErr(msg, err string, code int, causes []Causes) *Err {
	return &Err{
		Message: msg,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func (r *Err) Error() string {
	return r.Message
}

func NewBadRequestError(msg string) *Err {
	return &Err{
		Message: msg,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(msg string, causes []Causes) *Err {
	return &Err{
		Message: msg,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(msg string) *Err {
	return &Err{
		Message: msg,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(msg string) *Err {
	return &Err{
		Message: msg,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(msg string) *Err {
	return &Err{
		Message: msg,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
