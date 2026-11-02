package curr_errors

type CurrencyAlreadyExistsError struct {
	Code    string
	Massage string
}

func (e *CurrencyAlreadyExistsError) Error() string {
	return e.Massage + e.Code
}

func NewCurrencyAlreadyExistsError(code string) *CurrencyAlreadyExistsError {
	return &CurrencyAlreadyExistsError{Code: code, Massage: " - currency is already exists"}
}
