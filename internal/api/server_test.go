package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cryptomonitor/internal/api"
	"cryptomonitor/internal/domain"
)

func TestGETPrices(t *testing.T) {
	exchange := api.StubExchange{
		map[string]domain.Price{
			"bitcoin":  domain.Price{Coin: "bitcoin", Currency: "mxn", Value: 1309186},
			"ethereum": domain.Price{Coin: "ethereum", Currency: "mxn", Value: 40815},
		},
	}
	server := api.NewCryptoMonitorServer(&exchange)

	t.Run("regresar el precio de bitcoin en formato JSON", func(t *testing.T) {
		request := api.NewGetPriceRequest("bitcoin")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		api.AssertStatus(t, response.Code, http.StatusOK)

		var got domain.Price

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("no se pudo convertir la respuesta del servidor '%q' a una estructa Price, '%v'", response.Body, err)
		}
	})

	t.Run("regresar el precio de ethereum en formato JSON", func(t *testing.T) {
		request := api.NewGetPriceRequest("ethereum")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		api.AssertStatus(t, response.Code, http.StatusOK)

		var got domain.Price

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("no se pudo convertir la respuesta del servidor '%q' a una estructa Price, '%v'", response.Body, err)
		}
	})

	t.Run("regresa 404 con monedas no encontradas", func(t *testing.T) {
		request := api.NewGetPriceRequest("zzz")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got struct {
			ErrMsg string `json:"error"`
		}

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("no se pudo convertir la respuesta del servidor '%q', '%v'", response.Body, err)
		}

		api.AssertStatus(t, response.Code, http.StatusNotFound)
	})
}
