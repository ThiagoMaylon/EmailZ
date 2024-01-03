package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {

	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		println("nenhum erro")
	}
	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	fild := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(fild + " is required")
	case "min":
		return errors.New(fild + " is required min " + validationError.Param())
	case "max":
		return errors.New(fild + " is required max " + validationError.Param())
	case "email":
		return errors.New(fild + " is invalid ")
	}
	return nil
}
