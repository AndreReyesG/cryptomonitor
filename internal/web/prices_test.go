package web_test

import (
	"testing"
	"time"

	"cryptomonitor/internal/assert"
	"cryptomonitor/internal/domain"
	platformhttp "cryptomonitor/internal/platform/http"
	"cryptomonitor/internal/web"
)

var (
	fixedTime     = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	dummyStrSlice = []string{}
)

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
			stubClient := platformhttp.NewStub(200, tt.body)
			server, _ := web.NewCryptoMonitorServer(stubClient, dummyStrSlice)
			got, err := server.GetPrice(tt.coin)
			assert.NoError(t, err)
			assert.Price(t, got, tt.expectedPrice)
		})
	}
}
