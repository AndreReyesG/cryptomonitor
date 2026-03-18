package ui_test

import (
	"bytes"
	"testing"
	"time"

	"cryptomonitor/internal/domain"
	"cryptomonitor/ui"

	"github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	coins := []domain.Price{
		{
			Coin:        "bitcoin",
			Value:       1309186,
			LastUpdated: time.Now(),
			Exchange:    "coingecko",
		},
		{
			Coin:        "ethereum",
			Value:       40815,
			LastUpdated: time.Now(),
			Exchange:    "coingecko",
		},
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
