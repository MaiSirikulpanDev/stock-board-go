package models

import "time"

type Stock struct {
	Ticker     string    `json:"ticker"`
	Price      float64   `json:"price"`
	LastUpdate time.Time `json:"last_update"`
}
