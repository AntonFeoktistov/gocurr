package repository

import (
	"database/sql"
	"fmt"
)

type CurrencyRepo struct {
	db *sql.DB
}

func NewCurrencyRepo(db *sql.DB) *CurrencyRepo {
	return &CurrencyRepo{db: db}
}

func (repo *CurrencyRepo) GetByCode(code string) (*Currency, error) {
	var currency Currency
	row := repo.db.QueryRow(`SELECT Id, Code, FullName, Sign FROM Currencies WHERE Code = ?`, code)
	err := row.Scan(
		&currency.ID, &currency.Code, &currency.Name, &currency.Sign,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("currency not found: %s", code)
	}
	if err != nil {
		return nil, err
	}
	return &currency, nil
}

func (repo *CurrencyRepo) GetAll() ([]*Currency, error) {
	rows, err := repo.db.Query(`SELECT Id, Code, FullName, Sign FROM Currencies`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []*Currency
	for rows.Next() {
		var currency Currency
		err := rows.Scan(
			&currency.ID,
			&currency.Code,
			&currency.Name,
			&currency.Sign,
		)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, &currency)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return currencies, nil
}
