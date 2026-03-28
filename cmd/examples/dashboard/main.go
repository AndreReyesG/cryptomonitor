package main

import (
	"log"
	"net/http"
	"time"

	"cryptomonitor/internal/domain"
	"cryptomonitor/ui"
)

var (
	coingeckoURL = "https://www.coingecko.com/"
	prices       = []domain.Price{
		{
			Coin:        "bitcoin",
			Value:       1309186,
			LastUpdated: time.Now(),
			Exchange:    "coingecko",
			ExchangeURL: coingeckoURL,
		},
		{
			Coin:        "ethereum",
			Value:       40815,
			LastUpdated: time.Now(),
			Exchange:    "coingecko",
			ExchangeURL: coingeckoURL,
		},
	}
)

type application struct {
	prices []domain.Price
	templ  *ui.Renderer
}

func main() {
	app := &application{
		prices: prices,
	}

	templ, err := ui.NewRenderer()
	if err != nil {
		log.Fatal(err)
	}
	app.templ = templ

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.HandleFunc("GET /{$}", app.showDashboardHandler)
	mux.HandleFunc("POST /refresh", app.refreshHandler)

	log.Print("Empezando servidor en el puerto :5000")

	err = http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
