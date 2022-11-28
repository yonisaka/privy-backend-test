package helpers

import (
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Field " + fe.Field() + " is required"
	case "max":
		return "Should be less than " + fe.Param() + " characters"
	case "len":
		return "Should be grater than equal " + fe.Param() + " pajak"
	case "numeric":
		return "Should be numeric"
	case "gte":
		return "Should be greater than or equal to " + fe.Param()
	case "lte":
		return "Should be less than or equal to " + fe.Param()
	}
	return "Unknown error"
}
