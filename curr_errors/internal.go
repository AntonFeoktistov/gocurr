package curr_errors

type InternalError struct {
	Code    string
	Massage string
}

func (e *InternalError) Error() string {
	return e.Massage + e.Code
}

func NewInternalError(code string) *InternalError {
	return &InternalError{Code: code, Massage: " - Internal Error"}
}
