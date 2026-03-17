package domain

import "time"

type Price struct {
	Coin        string    `json:"coin"`
	Currency    string    `json:"currency"`
	Value       float64   `json:"value"`
	Exchange    string    `json:"exchange"`
	LastUpdated time.Time `json:"last_updated"`
}
