package core

import (
	"github.com/goexl/validate"
	"github.com/goexl/validator"
)

type Validator = validate.Validator

func newValidator() Validator {
	return validator.New()
}
