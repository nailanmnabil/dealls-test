package pkg

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func ExtendErr(custErr CustomError, err error) CustomError {
	custErr.Message = fmt.Sprintf("%s %s", custErr.Message, err.Error())
	return custErr
}

var (
	ErrInternal               = CustomError{Code: "E-00000", Message: "Sorry, we are experiencing issues, please try again later.", StatusCode: http.StatusInternalServerError}
	ErrResourceNotFound       = CustomError{Code: "E-00001", Message: "The data you are looking for was not found.", StatusCode: http.StatusNotFound}
	ErrBadRequest             = CustomError{Code: "E-00002", Message: "There is an error in your request.", StatusCode: http.StatusBadRequest}
	ErrUnauthorized           = CustomError{Code: "E-00003", Message: "You are not permitted to perform this action.", StatusCode: http.StatusUnauthorized}
	ErrForbidden              = CustomError{Code: "E-00004", Message: "You do not have the right to access this data.", StatusCode: http.StatusForbidden}
	ErrLimitReached           = CustomError{Code: "E-00005", Message: "Swipe limit for today has been reached.", StatusCode: http.StatusForbidden}
	ErrNoUserLeft             = CustomError{Code: "E-00006", Message: "No user left.", StatusCode: http.StatusNoContent}
	ErrProfileAlreadySwiped   = CustomError{Code: "E-00007", Message: "Profile already swiped.", StatusCode: http.StatusBadRequest}
	ErrActivePackageExist     = CustomError{Code: "E-00008", Message: "There is same active package exist.", StatusCode: http.StatusBadRequest}
	ErrEmailAlreadyRegistered = CustomError{Code: "E-00009", Message: "Email already registered.", StatusCode: http.StatusBadRequest}
	ErrInvalidEmailOrPass     = CustomError{Code: "E-00010", Message: "Invalid email or password.", StatusCode: http.StatusUnauthorized}
)
