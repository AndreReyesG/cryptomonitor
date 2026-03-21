package main

import (
	"log"
	"net/http"

	"cryptomonitor/internal/api"
	"cryptomonitor/internal/exchanges"
)

func main() {
	coingecko := exchanges.NewCoinGecko(exchanges.CoinGeckoAPIKey, http.DefaultClient, exchanges.RealTime{})
	server := api.NewCryptoMonitorServer(coingecko)
	log.Print("Iniciando servidor en el puerto :4000")
	log.Fatal(http.ListenAndServe(":4000", server))
}
