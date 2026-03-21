package exchanges_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"cryptomonitor/internal/exchanges"
)

const mockResponse = `{"bitcoin":{"usd":70737}}`

func TestCoinGeckoGetPrice(t *testing.T) {

	t.Run("PV-01: conectar y recibir respuesta de CoinGecko", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockResponse))
		}))
		defer server.Close()

		client := exchanges.NewCoinGeckoClient(server.URL)
		_, err := client.GetPrice("bitcoin")
		if err != nil {
			t.Errorf("PV-01 fallida: no se pudo conectar: %v", err)
		}
	})

	t.Run("PV-02: obtener precio de BTC y ETH correctamente", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockResponse))
		}))
		defer server.Close()

		client := exchanges.NewCoinGeckoClient(server.URL)
		price, err := client.GetPrice("bitcoin")
		if err != nil {
			t.Errorf("PV-02 fallida: error obteniendo BTC: %v", err)
		}
		if price.Value <= 0 {
			t.Errorf("PV-02 fallida: precio de BTC inválido: %v", price.Value)
		}
	})

	t.Run("PV-03: datos convertidos a estructura Price correctamente", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockResponse))
		}))
		defer server.Close()

		client := exchanges.NewCoinGeckoClient(server.URL)
		price, err := client.GetPrice("bitcoin")
		if err != nil {
			t.Fatalf("PV-03 fallida: error inesperado: %v", err)
		}
		if price.Coin != "bitcoin" {
			t.Errorf("PV-03 fallida: Coin esperado 'bitcoin', se obtuvo '%s'", price.Coin)
		}
		if price.Currency != "usd" {
			t.Errorf("PV-03 fallida: Currency esperado 'usd', se obtuvo '%s'", price.Currency)
		}
		if price.Exchange != "coingecko" {
			t.Errorf("PV-03 fallida: Exchange esperado 'coingecko', se obtuvo '%s'", price.Exchange)
		}
		if price.LastUpdated.IsZero() {
			t.Errorf("PV-03 fallida: LastUpdated no debe estar vacío")
		}
	})

	t.Run("PV-04: manejar error de conexión correctamente", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		server.Close()

		client := exchanges.NewCoinGeckoClient(server.URL)
		_, err := client.GetPrice("bitcoin")
		if err == nil {
			t.Errorf("PV-04 fallida: se esperaba un error pero no hubo ninguno")
		} else {
			t.Logf("PV-04: error manejado correctamente: %v", err)
		}
	})
}