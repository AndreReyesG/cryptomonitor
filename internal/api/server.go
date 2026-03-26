package api

import (
	"log/slog"
	"net/http"

	"cryptomonitor/internal/domain"
)

type CryptoMonitorServer struct {
	exchange domain.Exchange
	logger   *slog.Logger
	http.Handler
}

func NewCryptoMonitorServer(e domain.Exchange, l *slog.Logger) *CryptoMonitorServer {
	c := new(CryptoMonitorServer)

	c.exchange = e
	c.logger = l

	router := http.NewServeMux()
	router.Handle("/v1/prices/", http.HandlerFunc(c.showPriceHandler))

	c.Handler = router

	return c
}
