package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия соединения: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ошибка проверки соединения: %w", err)
	}

	if err := createTables(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("ошибка создания таблиц: %w", err)
	}

	log.Println("База данных и таблицы успешно инициализированы")
	return db, nil
}

func createTables(db *sql.DB) error {
	currenciesSQL := `
	CREATE TABLE IF NOT EXISTS Currencies (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Code VARCHAR(3) NOT NULL UNIQUE,
		FullName TEXT NOT NULL UNIQUE,
		Sign VARCHAR(3)
	);
	`

	exchangeRatesSQL := `
	CREATE TABLE IF NOT EXISTS ExchangeRates (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		BaseCurrencyId INTEGER NOT NULL,
		TargetCurrencyId INTEGER NOT NULL,
		Rate REAL NOT NULL,
		FOREIGN KEY (BaseCurrencyId) REFERENCES Currencies(Id) ON DELETE CASCADE,
		FOREIGN KEY (TargetCurrencyId) REFERENCES Currencies(Id) ON DELETE CASCADE
	);
	`

	_, err := db.Exec(currenciesSQL)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы Currencies: %w", err)
	}

	_, err = db.Exec(exchangeRatesSQL)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы ExchangeRates: %w", err)
	}

	return nil
}
