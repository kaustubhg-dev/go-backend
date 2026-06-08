package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

// ValidateStruct validates a struct and returns field → message map.
func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	result := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		result[e.Field()] = e.Tag()
	}

	return result
}