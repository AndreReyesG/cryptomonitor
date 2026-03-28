package main

import (
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"cryptomonitor/internal/domain"
)

func (app *application) showDashboardHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.templ.Render(w, app.prices, "dashboard.html"); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) refreshHandler(w http.ResponseWriter, r *http.Request) {
	app.prices = []domain.Price{
		{
			Coin:        "bitcoin",
			Value:       rand.Float64(),
			LastUpdated: time.Now(),
			Exchange:    "coingecko",
			ExchangeURL: coingeckoURL,
		},
		{
			Coin:        "ethereum",
			Value:       rand.Float64(),
			LastUpdated: time.Now(),
			Exchange:    "coingecko",
			ExchangeURL: coingeckoURL,
		},
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
