package exchanges

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cryptomonitor/internal/domain"
)

type CoinGeckoClient struct {
	BaseURL string
}

func NewCoinGeckoClient(baseURL string) *CoinGeckoClient {
	if baseURL == "" {
		baseURL = "https://api.coingecko.com"
	}
	return &CoinGeckoClient{BaseURL: baseURL}
}

type coinGeckoResponse map[string]map[string]interface{}

func (c *CoinGeckoClient) GetPrice(coin string) (domain.Price, error) {
	url := fmt.Sprintf(
		"%s/api/v3/simple/price?ids=%s&vs_currencies=usd",
		c.BaseURL,
		coin,
	)

	resp, err := http.Get(url)
	if err != nil {
		return domain.Price{}, fmt.Errorf("error consultando CoinGecko: %s", err.Error())
	}
	defer resp.Body.Close()

	var result coinGeckoResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return domain.Price{}, fmt.Errorf("error descodificando respuesta: %s", err.Error())
	}

	raw, ok := result[coin]["usd"]
	if !ok {
		return domain.Price{}, fmt.Errorf("no se encontró el precio para: %s", coin)
	}

	value, ok := raw.(float64)
	if !ok {
		return domain.Price{}, fmt.Errorf("formato de precio inválido para: %s", coin)
	}

	return domain.Price{
		Coin:        coin,
		Currency:    "usd",
		Value:       value,
		Exchange:    "coingecko",
		LastUpdated: time.Now(),
	}, nil
}