package api

import (
	"errors"
	"fmt"
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

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("no se obtuvo el estatus esperado, se obtuvo %d; se quizo %d", got, want)
	}
}
