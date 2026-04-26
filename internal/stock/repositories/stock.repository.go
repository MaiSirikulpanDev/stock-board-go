package repositories

import (
	"stock-board-go/internal/stock/models"
	"time"
)

type StockRepository interface {
	GetStock(ticker string, price float64) (*models.Stock, error)
}

type stockRepository struct{}

func NewStockRepository() StockRepository {
	return &stockRepository{}
}

func (r *stockRepository) GetStock(ticker string, price float64) (*models.Stock, error) {
	return &models.Stock{
		Ticker:     ticker,
		Price:      price,
		LastUpdate: time.Now(),
	}, nil
}
