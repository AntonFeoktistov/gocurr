package service

import "gocurr/repository"

type CurrencyService struct {
	repo *repository.CurrencyRepo
}

func NewCurrencyService(repo *repository.CurrencyRepo) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (service *CurrencyService) ListCurrencies() ([]*repository.Currency, error) {
	return service.repo.GetAll()
}
