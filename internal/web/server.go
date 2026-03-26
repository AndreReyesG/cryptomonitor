package web

import (
	"log/slog"
	"net/http"

	"cryptomonitor/internal/domain"
	platformhttp "cryptomonitor/internal/platform/http"
	"cryptomonitor/ui"
)

type CryptoMonitorServer struct {
	client platformhttp.Client
	coins  []string
	prices []domain.Price
	templ  *ui.Renderer
	logger *slog.Logger
	http.Handler
}

// NOTE: Tiene muchas dependencias.
// ¿Necesita un Builder?
func NewCryptoMonitorServer(
	client platformhttp.Client, coins []string, logger *slog.Logger,
) (*CryptoMonitorServer, error) {
	server := &CryptoMonitorServer{
		client: client,
		coins:  coins,
		logger: logger,
	}

	templ, err := ui.NewRenderer()
	if err != nil {
		return nil, err
	}

	server.templ = templ
	// TODO: Buscar mejor solución para el método GetPrices().
	server.prices = server.GetPrices()
	server.Handler = server.routes()

	return server, nil
}
