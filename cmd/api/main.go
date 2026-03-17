package main

import (
	"log"
	"net/http"

	"cryptomonitor/internal/api"
	"cryptomonitor/internal/domain"
)

type InMemoryExchange struct{}

func (i *InMemoryExchange) GetPrice(coin string) (domain.Price, error) {
	return domain.Price{
		Coin:     "dogecoin",
		Currency: "mxn",
		Value:    1.81,
	}, nil
}

func main() {
	server := api.NewCryptoMonitorServer(&InMemoryExchange{})
	log.Print("Iniciando servidor en el puerto :4000")
	log.Fatal(http.ListenAndServe(":4000", server))
}
