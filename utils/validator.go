package utils

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *CustomValidator) ValidateStruct(obj any) error {
	v.lazyinit()
	if errors, ok := v.validate.Struct(obj).(validator.ValidationErrors); ok {
		return NewValidationErrorsCustom(errors)
	}
	return nil
}

func (v *CustomValidator) Engine() any {
	v.lazyinit()
	return v.validate
}

func (v *CustomValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}
