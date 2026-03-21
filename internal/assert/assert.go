package assert

import (
	"errors"
	"net/http/httptest"
	"testing"

	"cryptomonitor/internal/domain"
)

func Error(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("se quiere un error pero no se obtiene uno")
	}

	if !errors.Is(got, want) {
		t.Errorf("se obtuvo '%v'; se queria '%v'", got, want)
	}
}

func NoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("se obtuvo un error, '%q'", got.Error())
	}
}

func Price(t testing.TB, got, want domain.Price) {
	t.Helper()
	if got.Coin != want.Coin {
		t.Fatalf("no concuerda domain.Price.Coin; se obtuvo %q, se queria %q", got.Coin, want.Coin)
	}

	if got.Currency != want.Currency {
		t.Fatalf("no concuerda domain.Price.Currency; se obtuvo %q, se queria %q",
			got.Currency, want.Currency)
	}

	if got.Value != want.Value {
		t.Fatalf("no concuerda domain.Price.Value; se obtuvo %f, se queria %f",
			got.Value, want.Value)
	}

	if got.Exchange != want.Exchange {
		t.Fatalf("no concuerda domain.Price.Exchange; se obtuvo %q, se queria %q",
			got.Exchange, want.Exchange)
	}

	if got.LastUpdated != want.LastUpdated {
		t.Fatalf("no concuerda domain.Price.LastUpdated; se obtuvo %v, se queria %v",
			got.LastUpdated, want.LastUpdated)
	}
}

func Status(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("no se obtuvo el estatus esperado; se obtuvo %d, se queria %d", got, want)
	}
}

func ContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("la respuesta no tenia content-type de tipo %v, se obtuvo %v",
			want, response.Result().Header)
	}
}
