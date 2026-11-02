package service

import (
	"errors"
	"gocurr/curr_errors"
	"gocurr/repository"
	"log"
	"strings"
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

func (service *CurrencyService) CreateCurrency(name, code, sign string) (*repository.Currency, error) {

	if !validateCode(code) || !validateName(name) || !validateSign(sign) {
		err := curr_errors.NewNotValidError(code)
		return nil, err
	}
	name = strings.ToUpper(name)
	code = strings.ToUpper(code)
	sign = strings.ToUpper(sign)

	currency, err := service.repo.GetByCode(code)
	if errors.As(err, new(*curr_errors.InternalError)) {
		return nil, err
	}
	if currency != nil {
		return nil, curr_errors.NewCurrencyAlreadyExistsError(code)
	}
	newCurrency, err := service.repo.CreateCurrency(name, code, sign)
	log.Printf("Service: %s", newCurrency)
	if err != nil {
		return nil, err
	}
	return newCurrency, nil
}
