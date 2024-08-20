package repositories

import "errors"

var (
	ErrDuplicatedKeyEmail       = errors.New("duplicated email not allowed")
	ErrDuplicatedKeyPhoneNumber = errors.New("duplicated phone_number not allowed")
)
