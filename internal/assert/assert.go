package assert

import (
	"errors"
	"testing"
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
