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

// TODO: Buscar una mejor solución para el método GetPrices.
// TODO: Agregar errores.
func (s *CryptoMonitorServer) GetPrices() []domain.Price {
	var prices []domain.Price

	for _, coin := range s.coins {
		p, err := s.GetPrice(coin)
		if err != nil {
			continue
		}
		prices = append(prices, p)
	}

	return prices
}
