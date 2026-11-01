package main

import (
	"context"
	"errors"
	"gocurr/database"
	"gocurr/handlers"
	"gocurr/repository"
	"gocurr/service"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	db, err := database.InitDB("./database/currdb.sqlite")
	if err != nil {
		log.Fatalf("Не удалось инициализировать БД: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	currencyRepo := repository.NewCurrencyRepo(db)
	currencyService := service.NewCurrencyService(currencyRepo)
	currencyHandler := handlers.NewCurrencyHandler(currencyService)

	mux.HandleFunc("GET /currencies", currencyHandler.List)
	mux.HandleFunc("GET /currency/{code}", currencyHandler.GetByCode)
	mux.HandleFunc("GET /currency/", func(w http.ResponseWriter, r *http.Request) {
		handlers.WriteJSONError(w, http.StatusBadRequest, "Код валюты отсутствует в адресе")
	})
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           loggingMiddleware(mux),
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	go func() {
		slog.Info("server started", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server error", "err", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	slog.Info("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("shutdown error", "err", err)
	} else {
		slog.Info("server stopped gracefully")
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
	})
}
