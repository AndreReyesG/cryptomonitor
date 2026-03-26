package web_test

import (
	"log/slog"
	"testing"
	"time"

	"cryptomonitor/internal/assert"
	"cryptomonitor/internal/domain"
	phttp "cryptomonitor/internal/platform/http"
	"cryptomonitor/internal/web"
)

var (
	fixedTime = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
)

// NOTE: ¿Esto es un Proxy?
func newTestServer(t *testing.T, c phttp.Client) *web.CryptoMonitorServer {
	logger := slog.New(slog.DiscardHandler)
	server, err := web.NewCryptoMonitorServer(c, []string{}, logger)
	if err != nil {
		t.Fatalf("problema creado el servidor, %q", err.Error())
	}
	return server
}

const (
	btcJSON = `{"coin":"bitcoin","currency":"mxn","value":1264773,"exchange":"coingecko","last_updated":"2026-01-01T00:00:00Z"}`
	ethJSON = `{"coin":"ethereum","currency":"mxn","value":38538,"exchange":"coingecko","last_updated":"2026-01-01T00:00:00Z"}`
)

func TestGETPrice(t *testing.T) {
	tests := []struct {
		name          string
		coin          string
		body          string
		expectedPrice domain.Price
	}{
		{
			name: "obtener el precio de bitcoin desde nuestra API interna",
			coin: "bitcoin",
			body: btcJSON,
			expectedPrice: domain.Price{
				Coin:        "bitcoin",
				Currency:    "mxn",
				Value:       1264773,
				Exchange:    "coingecko",
				LastUpdated: fixedTime,
			},
		},
		{
			name: "obtener el precio de ethereum desde nuestra API interna",
			coin: "ethereum",
			body: ethJSON,
			expectedPrice: domain.Price{
				Coin:        "ethereum",
				Currency:    "mxn",
				Value:       38538,
				Exchange:    "coingecko",
				LastUpdated: fixedTime,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stubClient := phttp.NewStub(200, tt.body)
			server := newTestServer(t, stubClient)
			got, err := server.GetPrice(tt.coin)
			assert.NoError(t, err)
			assert.Price(t, got, tt.expectedPrice)
		})
	}
}
