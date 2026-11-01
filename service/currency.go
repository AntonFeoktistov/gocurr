package service

import (
	"gocurr/curr_errors"
	"gocurr/repository"
)

type CurrencyService struct {
	repo *repository.CurrencyRepo
}

func NewCurrencyService(repo *repository.CurrencyRepo) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (service *CurrencyService) ListCurrencies() ([]*repository.Currency, error) {
	return service.repo.GetAll()
}

func (service *CurrencyService) GetByCode(code string) (*repository.Currency, error) {
	if !validateCode(code) {
		err := curr_errors.NewNotValidError(code)
		return nil, err
	}
	currency, err := service.repo.GetByCode(code)

	if err != nil {
		return nil, err
	}
	return currency, nil
}
