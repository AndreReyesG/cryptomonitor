package web

import (
	"encoding/json"
	"fmt"

	"cryptomonitor/internal/domain"
)

func (s *CryptoMonitorServer) GetPrice(coin string) (domain.Price, error) {
	url := fmt.Sprintf("http://localhost:4000/v1/prices/%s", coin)
	resp, err := s.client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return domain.Price{}, fmt.Errorf("cryptomonitor: error al hacer request: %w", err)
	}

	var p domain.Price
	err = json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// TODO: Buscar una mejor soloción para el método GetPrices ¿Un Stub?
// TODO: Agregar logs.
func (s *CryptoMonitorServer) GetPrices(coins []string) []domain.Price {
	var prices []domain.Price

	for _, coin := range coins {
		p, err := s.GetPrice(coin)
		if err != nil {
			continue
		}
		prices = append(prices, p)
	}

	return prices
}
