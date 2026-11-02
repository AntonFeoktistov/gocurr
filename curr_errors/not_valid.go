package curr_errors

type NotValidError struct {
	Code    string
	Massage string
}

func (e *NotValidError) Error() string {
	return e.Massage + e.Code
}

func NewNotValidError(code string) *NotValidError {
	return &NotValidError{Code: code, Massage: " - not valid data"}
}
