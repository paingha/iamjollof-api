package utils

import (
	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(structs interface{}) error {
	err := validate.Struct(structs)
	if err != nil {
		return err
	}
	return nil
}

func ValidateVariable(vars interface{}, rules string) error {
	err := validate.Var(vars, rules) //rules: "required,email"
	if err != nil {
		return err
	}
	return nil
}
