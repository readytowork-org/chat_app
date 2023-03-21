package validators

import (
	"fmt"
	"letschat/errors"
	"regexp"

	validator "github.com/go-playground/validator/v10"
)

// UserValidator structure
type UserValidator struct {
	Validate *validator.Validate
}

// Register Custom Validators
func NewUserValidator() UserValidator {
	v := validator.New()
	_ = v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		if fl.Field().String() != "" {
			match, _ := regexp.MatchString("^[0-9]{10}$", fl.Field().String())
			return match
		}
		return true
	})
	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		if fl.Field().String() != "" {
			password := fl.Field().String()
			if len(password) < 6 {
				return false
			}
		}
		return true
	})
	_ = v.RegisterValidation("confirm_password", func(fl validator.FieldLevel) bool {
		if fl.Field().String() != "" {
			password := fl.Field().String()
			if len(password) < 6 {
				return false
			}
		}
		return true
	})

	return UserValidator{
		Validate: v,
	}
}

func (cv UserValidator) generateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	case "password":
		return fmt.Sprintf("Field '%s' is not valid.", field)
	case "confirm_password":
		return fmt.Sprintf("Field '%s' is not valid.", field)
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}

func (cv UserValidator) GenerateValidationResponse(err error) []errors.ErrorContext {
	var validations []errors.ErrorContext
	for _, value := range err.(validator.ValidationErrors) {
		field, rule := value.Field(), value.Tag()
		validation := errors.ErrorContext{Field: field, Message: cv.generateValidationMessage(field, rule)}
		validations = append(validations, validation)
	}
	return validations
}
