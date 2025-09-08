package validation

import (
	"fmt"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func Init() {
	Validate.RegisterValidation("strong_password", validatePassword)
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var (
		hasUpper   = false
		hasLower   = false
		hasDigit   = false
		hasSpecial = false
	)

	if len(password) < 8 {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsNumber(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}

func FormatValidationErrors(err error) map[string]string {
	errorMessages := make(map[string]string)
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			tag := fieldError.Tag()

			switch tag {
			case "required":
				errorMessages[fieldName] = fmt.Sprintf("The %s field is required.", fieldName)
			case "min":
				errorMessages[fieldName] = fmt.Sprintf("The %s field must be at least %s characters long.", fieldName, fieldError.Param())
			case "max":
				errorMessages[fieldName] = fmt.Sprintf("The %s field must be at most %s characters long.", fieldName, fieldError.Param())
			case "email":
				errorMessages[fieldName] = "The email format is invalid."
			case "strong_password":
				errorMessages[fieldName] = "The password must contain at least one uppercase letter, one lowercase letter, one digit, and one special character."
			default:
				errorMessages[fieldName] = fmt.Sprintf("The %s field is not valid.", fieldName)
			}
		}
	}
	return errorMessages
}
