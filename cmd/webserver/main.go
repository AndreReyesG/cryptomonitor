package main

import (
	"log"
	"net/http"
	"time"

	"cryptomonitor/internal/domain"
	"cryptomonitor/ui"
)

func main() {
	coins := []domain.Price{
		{
			Coin:        "bitcoin",
			Value:       1309186,
			LastUpdated: time.Now(),
			Exchange:    "coingecko",
		},
		{
			Coin:        "ethereum",
			Value:       40815,
			LastUpdated: time.Now(),
			Exchange:    "coingecko",
		},
	}

	dashboard, err := ui.NewRenderer()
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

	router.Handle("/{$}", http.HandlerFunc(func(w http.ResponseWriter, h *http.Request) {
		if err := dashboard.Render(w, coins, "dashboard.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

	log.Print("Iniciando servidor en el puerto :9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
