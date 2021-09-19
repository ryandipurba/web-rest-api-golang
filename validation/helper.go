package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errors []string
	var resultError []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	var join string
	for _, str := range errors {
		join += str
	}

	splited := strings.Split(join, ".")
	if splited[1] == "Name' Error:Field validation for 'Name' failed on the 'required' tag" {
		resultError = append(resultError, "Name must be input")
	} else if splited[1] == "Creator' Error:Field validation for 'Creator' failed on the 'required' tag" {
		resultError = append(resultError, "Creator must be input")
	} else {
		resultError = append(resultError, "Name and Creator must be input")
	}

	return resultError
}
