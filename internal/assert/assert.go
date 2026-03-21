package assert

import (
	"errors"
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
