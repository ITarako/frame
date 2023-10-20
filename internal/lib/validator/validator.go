package validator

import (
	"fmt"
	pkgValidator "github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type Validator struct {
	validate *pkgValidator.Validate
	Errors   map[string]string
}

func NewValidator() *Validator {
	validate := pkgValidator.New(pkgValidator.WithRequiredStructEnabled())

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return &Validator{
		validate: validate,
		Errors:   make(map[string]string),
	}
}

func (v *Validator) ValidateStruct(dest any) {
	if err := v.validate.Struct(dest); err != nil {
		if _, ok := err.(*pkgValidator.InvalidValidationError); ok {
			panic(err)
		}

		validateErrors := err.(pkgValidator.ValidationErrors)
		v.parseErrors(validateErrors)
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) parseErrors(validateErrors pkgValidator.ValidationErrors) {
	for _, err := range validateErrors {
		var message string
		field := err.Field()

		switch err.ActualTag() {
		case "required":
			message = "is required"
		case "len":
			message = fmt.Sprintf("is should be length %s", err.Param())
		case "min":
			message = fmt.Sprintf("min length %s", err.Param())
		case "max":
			message = fmt.Sprintf("max length %s", err.Param())
		case "url":
			message = "is not a valid URL"
		case "email":
			message = "is not a valid email"
		case "e164":
			message = "is not a valid phone number"
		default:
			message = "is not valid"
		}

		v.AddError(field, message)
	}
}
