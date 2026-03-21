package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"cryptomonitor/internal/domain"
)

type StubExchange struct {
	Coins map[string]domain.Price
}

func (s *StubExchange) GetPrice(coin string) (domain.Price, error) {
	if coin != "bitcoin" && coin != "ethereum" {
		return domain.Price{}, errors.New("oh no")
	}
	return s.Coins[coin], nil
}

func NewGetPriceRequest(coin string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/prices/%s", coin), nil)
	return req
}

func GetPriceFromResponse(t testing.TB, body io.Reader) (price domain.Price) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&price)
	if err != nil {
		t.Fatalf("no se pudo convertir la respuesta del servidor %q, '%v'", body, err)
	}
	return
}
