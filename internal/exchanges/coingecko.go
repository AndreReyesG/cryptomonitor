package exchanges

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cryptomonitor/internal/domain"
)

type CoinGeckoClient struct{}

type coinGeckoResponse map[string]map[string]float64

func (c *CoinGeckoClient) GetPrice(coin string) (domain.Price, error) {
	url:=fmt.Sprintf(
		"https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd",
		coin,
	)

	resp, err:=http.Get(url)
	if err!=nil {
		return domain.Price{}, fmt.Errorf("error consultando CoinGecko: %s", err.Error())
	}
	defer resp.Body.Close()

	var result coinGeckoResponse
	err=json.NewDecoder(resp.Body).Decode(&result)
	if err!=nil {
		return domain.Price{}, fmt.Errorf("error descodificando respuesta: %s", err.Error())
	}

	value, ok:=result[coin]["usd"]
	if !ok {
		return domain.Price{}, fmt.Errorf("no se encontró el precio para: %s", coin)
	}

	return domain.Price{
		Coin:        coin,
		Currency:    "usd",
		Value:       value,
		Exchange:    "coingecko",
		LastUpdated: time.Now(),
	}, nil
}