package exceptions

import (
	"privy-backend-test/internal/helpers"
)

type Error struct {
	Code    int
	Message interface{}
	Status  int
}

func ErrorException(code int, msg string) Error {
	return Error{
		Code:    code,
		Message: msg,
		Status:  406,
	}
}

func ErrorValidationException(code int, msg []helpers.ErrorMsg) Error {
	return Error{
		Code:    code,
		Message: msg,
		Status:  406,
	}
}

func AuthException(code int, msg string) Error {
	return Error{
		Code:    code,
		Message: msg,
		Status:  401,
	}
}
