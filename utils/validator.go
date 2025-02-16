package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct(s any) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = "Invalid " + e.Field()
	}
	return errors
}
