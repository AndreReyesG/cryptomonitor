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

	dashboard, err := ui.NewRenderer()
	if err != nil {
		t.Fatal(err)
	}

	page := "dashboard.html"

	t.Run("muestra monedas en HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := dashboard.Render(&buf, coins, page); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
