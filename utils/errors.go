package utils

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var mapped = map[string]func(validator.FieldError) string{
	"required": func(fe validator.FieldError) string {
		return fmt.Sprintf("El campo %s no puede estar vacío", fe.Field())
	},
	"email": func(fe validator.FieldError) string {
		return fmt.Sprintf("El campo %s no es un email válido", fe.Field())
	},
	"len": func(fe validator.FieldError) string {
		return fmt.Sprintf("El campo %s no cumple con la longitud de %s", fe.Field(), fe.Param())
	},
	"number": func(fe validator.FieldError) string {
		return fmt.Sprintf("El campo %s no es un número válido", fe.Field())
	},
	"min": func(fe validator.FieldError) string {
		return fmt.Sprintf("El campo %s no cumple con la longitud mínima de %s", fe.Field(), fe.Param())
	},
}

type validationErrorsCustom struct {
}

func NewValidationErrorsCustom(err error) validationErrorsCustom {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return validationErrorsCustom{fieldErrors: validationErrors}
	}
	return validationErrorsCustom{}
}

func (c validationErrorsCustom) Error() string {
	buff := bytes.NewBufferString("")

	max := len(c.fieldErrors)
	for i := 0; i < max-1; i++ {
		buff.WriteString(messageBuild(c.fieldErrors[i]))
		buff.WriteString(";")
	}
	buff.WriteString(messageBuild(c.fieldErrors[max-1]))
	return strings.TrimSpace(buff.String())
}

func messageBuild(fe validator.FieldError) string {
	if fn, ok := mapped[fe.Tag()]; ok {
		return fn(fe)
	}
	return fe.Error()
}
