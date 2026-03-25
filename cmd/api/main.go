package main

import (
	"log"
	"net/http"

	"cryptomonitor/internal/api"
	"cryptomonitor/internal/exchanges"
)

func main() {
	coingecko := exchanges.NewCoinGecko(exchanges.CoinGeckoAPIKey, http.DefaultClient, nil)
	log.Print("Iniciando servidor en el puerto :4000")
	proxy := exchanges.NewCoinGeckoProxy(coingecko)
	server := api.NewCryptoMonitorServer(proxy)
	log.Print("Iniciando Proxy de CoinGecko...")
	log.Fatal(http.ListenAndServe(":4000", server))
}
