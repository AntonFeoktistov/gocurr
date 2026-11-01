package repository

type Currency struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Sign string `json:"sign,omitempty"`
}

type ExchangeRate struct {
	ID             int
	BaseCurrency   Currency
	TargetCurrency Currency
	Rate           float64
}
