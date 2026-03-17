package domain

type Exchange interface {
	GetPrice(coin string) (Price, error)
}
