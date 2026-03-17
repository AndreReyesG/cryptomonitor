package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"cryptomonitor/internal/domain"
)

type CryptoMonitorServer struct {
	exchange domain.Exchange
}

func NewCryptoMonitorServer(e domain.Exchange) *CryptoMonitorServer {
	return &CryptoMonitorServer{
		exchange: e,
	}
}

func (c *CryptoMonitorServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	coin := strings.TrimPrefix(r.URL.Path, "/v1/prices/")

	price, err := c.exchange.GetPrice(coin)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		msg := struct {
			ErrMsg string `json:"error"`
		}{
			ErrMsg: "oh no",
		}

		js, _ := json.Marshal(msg)
		w.Write(js)
		return
	}

	js, _ := json.Marshal(price)
	w.Write(js)
}
