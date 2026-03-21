package exchanges

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"cryptomonitor/internal/domain"
	phttp "cryptomonitor/internal/platform/http"
	ptime "cryptomonitor/internal/platform/time"
)

type CoinGecko struct {
	key    string
	client phttp.Client
	clock  ptime.Provider
}

func NewCoinGecko(key string, client phttp.Client, clock ptime.Provider) *CoinGecko {
	if clock == nil {
		clock = ptime.Real{}
	}

	return &CoinGecko{
		key:    key,
		client: client,
		clock:  clock,
	}
}

type CoinGeckoResponse map[string]map[string]float64

const CoinGeckoAPIKey = "CG-Xkt8ErqVvXxSxUUMBEds5Zhq"

var (
	ErrAPIKeyMissing = errors.New("coinGecko: API key invalida")
	ErrCoinNotFound  = errors.New("coinGecko: moneda no encontrada")
)

func (c *CoinGecko) GetPrice(coin string) (domain.Price, error) {
	url := fmt.Sprintf(
		"https://api.coingecko.com/api/v3/simple/price?vs_currencies=mxn&ids=%s&x_cg_demo_api_key=%s",
		coin,
		c.key,
	)

	resp, err := c.client.Get(url)
	if err != nil {
		return domain.Price{}, fmt.Errorf("coinGecko: error al hacer request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return domain.Price{}, ErrAPIKeyMissing
	}

	coinGeckoResponse := make(CoinGeckoResponse)
	err = json.NewDecoder(resp.Body).Decode(&coinGeckoResponse)
	if err != nil {
		return domain.Price{}, fmt.Errorf("coinGecko: error al decodificar respuesta: %w", err)
	}

	price, ok := coinGeckoResponse[coin]
	if !ok {
		return domain.Price{}, ErrCoinNotFound
	}

	return domain.Price{
		Coin:        coin,
		Currency:    "mxn",
		Exchange:    "coingecko",
		ExchangeURL: "https://www.coingecko.com/",
		Value:       price["mxn"],
		LastUpdated: c.clock.Now(),
	}, nil
}
