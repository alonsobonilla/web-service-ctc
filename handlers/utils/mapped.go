package handlers

import (
	"ctcwebservice/repositories"
	"net/http"
)

// Mapping of errors and status codes
var Mapped = map[error]int{
	repositories.ErrDuplicatedKeyEmail:            http.StatusBadRequest,
	repositories.ErrDuplicatedKeyPhoneNumber:      http.StatusBadRequest,
	repositories.ErrDuplicatedKeyEmailPhoneNumber: http.StatusBadRequest,
	repositories.ErrRecordNotFound:                http.StatusNotFound,
}
