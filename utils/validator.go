package utils

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *Validator) ValidateStruct(obj any) error {
	v.lazyinit()
	if err := v.validate.Struct(obj); err != nil {
		return NewCustomErrors(err)
	}
	return nil
}

func (v *Validator) Engine() any {
	v.lazyinit()
	return v.validate
}

func (v *Validator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}
