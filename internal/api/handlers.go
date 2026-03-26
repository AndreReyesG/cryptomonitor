package api

import (
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

		err = writeJSON(w, http.StatusNotFound, errMsg)
		if err != nil {
			c.logger.Error(err.Error())
			http.Error(w, "El servidor encontro un problema y no puede procesar tu request", http.StatusInternalServerError)
		}
		return
	}

	err = writeJSON(w, http.StatusOK, price)
	if err != nil {
		c.logger.Error(err.Error())
		http.Error(w, "El servidor encontro un problema y no puede procesar tu request", http.StatusInternalServerError)
	}
}
