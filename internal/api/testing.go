package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("no se obtuvo el estatus esperado; se obtuvo %d, se quizo %d", got, want)
	}
}

func AssertPrice(t testing.TB, got, want domain.Price) {
	t.Helper()
	if got.Coin != want.Coin {
		t.Fatalf("no concuerda domain.Price.Coin; se obtuvo %q, se quizo %q", got.Coin, want.Coin)
	}

	if got.Currency != want.Currency {
		t.Fatalf("no concuerda domain.Price.Currency; se obtuvo %q, se quizo %q",
			got.Currency, want.Currency)
	}

	if got.Value != want.Value {
		t.Fatalf("no concuerda domain.Price.Value; se obtuvo %f, se quizo %f",
			got.Value, want.Value)
	}

	if got.Exchange != want.Exchange {
		t.Fatalf("no concuerda domain.Price.Exchange; se obtuvo %q, se quizo %q",
			got.Exchange, want.Exchange)
	}

	if got.LastUpdated != want.LastUpdated {
		t.Fatalf("no concuerda domain.Price.LastUpdated; se obtuvo %v, se quizo %v",
			got.LastUpdated, want.LastUpdated)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("la respuesta no tenia content-type de tipo %v, se obtuvo %v",
			want, response.Result().Header)
	}
}
