package curr_errors

type NotFoundError struct {
	Code    string
	Massage string
}

func (e *NotFoundError) Error() string {
	return e.Massage + e.Code
}

func NewNotFoundError(code string) *NotFoundError {
	return &NotFoundError{Code: code, Massage: " - not found"}
}
