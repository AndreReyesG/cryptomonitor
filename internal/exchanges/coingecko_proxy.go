package exchanges

import (
	"cryptomonitor/internal/domain"
	"log"
)

// CoinGeckoProxy implementa domain.Exchange y delega al exchange real.
// Permite agregar comportamiento transversal (caché, logs, rate limiting)
// sin modificar CoinGecko ni los servidores que lo consumen.
type CoinGeckoProxy struct {
	real domain.Exchange
}

func NewCoinGeckoProxy(real domain.Exchange) *CoinGeckoProxy {
	return &CoinGeckoProxy{real: real}
}

func (p *CoinGeckoProxy) GetPrice(coin string) (domain.Price, error) {

	// lógica antes de delegar (caché, logs, rate limit...)
	log.Print("Realizando request a CoinGeckoProxy para la moneda: ", coin)
	price, err := p.real.GetPrice(coin)

	// POST: aquí irá lógica después (guardar en caché, registrar respuesta...)

	return price, err
}
