package repositories

import "errors"

var (
	ErrDuplicatedKeyEmail            = errors.New("ErrDuplicatedKeyEmail")
	ErrDuplicatedKeyPhoneNumber      = errors.New("ErrDuplicatedKeyPhoneNumber")
	ErrDuplicatedKeyEmailPhoneNumber = errors.New("ErrDuplicatedKeyEmailPhoneNumber")
	ErrRecordNotFound                = errors.New("ErrRecordNotFound")
)
