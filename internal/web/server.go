package web

import (
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
	http.Handler
}

// TODO: Necita un Builder.
// TODO: Agregar logs.
func NewCryptoMonitorServer(client platformhttp.Client, coins []string) (*CryptoMonitorServer, error) {
	c := &CryptoMonitorServer{
		client: client,
		coins:  coins,
	}

	templ, err := ui.NewRenderer()
	if err != nil {
		return nil, err
	}

	c.templ = templ

	// TODO: Buscar una mejor soloción para el método GetPrices ¿Un Stub?
	c.prices = c.GetPrices(coins)

	router := http.NewServeMux()
	router.Handle("/{$}", http.HandlerFunc(func(w http.ResponseWriter, h *http.Request) {
		if err := c.templ.Render(w, c.prices, "dashboard.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

	c.Handler = router

	return c, nil
}
