package repositories

import "errors"

var (
	ErrDuplicatedKeyEmail            = errors.New("correo ya registrado")
	ErrDuplicatedKeyPhoneNumber      = errors.New("número de teléfono ya registrado")
	ErrDuplicatedKeyEmailPhoneNumber = errors.New("número de teléfono y correo ya registrados")
)
