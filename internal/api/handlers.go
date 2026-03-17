package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (c *CryptoMonitorServer) showPriceHandler(w http.ResponseWriter, r *http.Request) {
	coin := strings.TrimPrefix(r.URL.Path, "/v1/prices/")

	price, err := c.exchange.GetPrice(coin)
	if err != nil {
		errMsg := map[string]string{
			"error": err.Error(),
		}

		js, err := json.Marshal(errMsg)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "El servidor encontro un problema y no puede procesar tu request", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(js)
		return
	}

	js, err := json.Marshal(price)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "El servidor encontro un problema y no puede procesar tu request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(js)
}
