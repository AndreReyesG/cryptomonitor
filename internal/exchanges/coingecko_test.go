package exchanges_test

import (
	"net/http"
	"testing"
	"time"

	"cryptomonitor/internal/assert"
	"cryptomonitor/internal/domain"
	"cryptomonitor/internal/exchanges"
	phttp "cryptomonitor/internal/platform/http"
	ptime "cryptomonitor/internal/platform/time"
)

var (
	fixedTime     = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	stubTime      = ptime.Stub{T: fixedTime}
	dummyStubTime = ptime.Stub{}
	coingeckoURL  = "https://www.coingecko.com/"
)

// NOTE: TestCoinGeckoExchange son pruebas de integración.
func TestCoinGeckoExchange(t *testing.T) {
	t.Skip("saltar pruebas de integracion")
	t.Run("regresa error cuando la api key no es valida", func(t *testing.T) {
		coingecko := exchanges.NewCoinGecko("pa55word", http.DefaultClient, nil)
		_, err := coingecko.GetPrice("bitcoin")
		assert.Error(t, err, exchanges.ErrAPIKeyMissing)
	})
}

func TestCoinGecko_ErrorAPIKey(t *testing.T) {
	t.Run("regresa error cuando la api key no es valida", func(t *testing.T) {
		stubClient := phttp.NewStub(http.StatusUnauthorized, `{}`)
		coingecko := exchanges.NewCoinGecko("pa55word", stubClient, dummyStubTime)
		_, err := coingecko.GetPrice("bitcoin")
		assert.Error(t, err, exchanges.ErrAPIKeyMissing)
	})
}

func TestCoinGecko_ErrorCoinNotFound(t *testing.T) {
	t.Run("regresa error con moneda inexistente", func(t *testing.T) {
		stubClient := phttp.NewStub(http.StatusOK, `{}`)
		coingecko := exchanges.NewCoinGecko("pa55word", stubClient, dummyStubTime)
		_, err := coingecko.GetPrice("chale")
		assert.Error(t, err, exchanges.ErrCoinNotFound)
	})
}

func TestCoinGecko_ReturnBitcoin(t *testing.T) {
	t.Run("regresar bitcoin", func(t *testing.T) {
		body := `{"bitcoin":{"mxn":1309186}}`
		stubClient := phttp.NewStub(http.StatusOK, body)
		coingecko := exchanges.NewCoinGecko("pa55word", stubClient, stubTime)

		got, _ := coingecko.GetPrice("bitcoin")
		want := domain.Price{
			Coin:        "bitcoin",
			Currency:    "mxn",
			Exchange:    "coingecko",
			ExchangeURL: coingeckoURL,
			Value:       1309186,
			LastUpdated: fixedTime,
		}

		assert.Price(t, got, want)
	})
}
