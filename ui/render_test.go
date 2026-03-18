package ui_test

import (
	"bytes"
	"testing"

	"cryptomonitor/internal/domain"
	"cryptomonitor/ui"

	"github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	coins := []domain.Price{
		{Coin: "bitcoin"},
		{Coin: "ethereum"},
	}

	t.Run("muestra monedas en HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := ui.Render(&buf, coins); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
