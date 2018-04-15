package resource

import (
	"net/http"
	"strconv"
)

// Status ...
type Status struct {
	Success bool
	Code    int
	Message string
}

// Success ...
func Success(code int) Status {
	return Status{
		Success: true,
		Code:    code,
	}
}

// Fail ...
func Fail(code int, message string) Status {
	return Status{
		Success: false,
		Code:    code,
		Message: message,
	}
}

// FailSimple ...
func FailSimple(code int) Status {
	return Status{
		Success: false,
		Code:    code,
		Message: strconv.Itoa(code) + " " + http.StatusText(code),
	}
}
