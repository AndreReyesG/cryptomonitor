package main

import (
	"log/slog"
	"net/http"
	"os"

	"cryptomonitor/internal/api"
	"cryptomonitor/internal/exchanges"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	coingecko := exchanges.NewCoinGecko(exchanges.CoinGeckoAPIKey, http.DefaultClient, nil)
	server := api.NewCryptoMonitorServer(coingecko)

	logger.Info("inicidando servidor", "addr", ":4000")

	err := http.ListenAndServe(":4000", server)
	logger.Error(err.Error())
	os.Exit(1)
}
