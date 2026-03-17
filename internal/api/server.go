package api

import (
	"net/http"

	"cryptomonitor/internal/domain"
)

type CryptoMonitorServer struct {
	exchange domain.Exchange
	http.Handler
}

func NewCryptoMonitorServer(e domain.Exchange) *CryptoMonitorServer {
	c := new(CryptoMonitorServer)

	c.exchange = e
	
	router := http.NewServeMux()
	router.Handle("/v1/prices/", http.HandlerFunc(c.showPriceHandler))

	c.Handler = router

	return c
}
