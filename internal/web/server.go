package web

import (
	"net/http"

	"cryptomonitor/internal/domain"
	platformhttp "cryptomonitor/internal/platform/http"
	"cryptomonitor/ui"
)

type CryptoMonitorServer struct {
	client platformhttp.Client
	coins  []string // NOTE: ¿Realmente se necesita?
	prices []domain.Price
	templ  *ui.Renderer
	http.Handler
}

// NOTE: ¿Necesita un Builder?
// TODO: Agregar logs.
func NewCryptoMonitorServer(client platformhttp.Client, coins []string) (*CryptoMonitorServer, error) {
	server := &CryptoMonitorServer{
		client: client,
		coins:  coins,
	}

	templ, err := ui.NewRenderer()
	if err != nil {
		return nil, err
	}

	server.templ = templ

	// TODO: Buscar una mejor soloción para el método GetPrices ¿Un Stub?
	server.prices = server.GetPrices(coins)

	router := http.NewServeMux()
	router.Handle("/{$}", http.HandlerFunc(server.showDashboardHandler))

	server.Handler = router

	return server, nil
}
