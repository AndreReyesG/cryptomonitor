package main

import (
	"log/slog"
	"net/http"
	"os"

	"cryptomonitor/internal/api"
	"cryptomonitor/internal/exchanges"
)

func main() {
	coingecko := exchanges.NewCoinGecko(exchanges.CoinGeckoAPIKey, http.DefaultClient, nil)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	proxy := exchanges.NewCoinGeckoProxy(coingecko)
	server := api.NewCryptoMonitorServer(proxy, logger)

	logger.Info("Iniciando Servidor", "addr", ":4000")
	logger.Info("Iniciando Proxy de CoinGecko...")

	err := http.ListenAndServe(":4000", server)
	logger.Error(err.Error())
	os.Exit(1)
}
