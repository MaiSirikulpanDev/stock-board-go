package repositories

import (
	"math/rand"
	"stock-board-go/internal/stock/models"
	"time"
)

type StockRepository interface {
	GetStock(ticker string) (*models.Stock, error)
}

type stockRepository struct{}

func NewStockRepository() StockRepository {
	return &stockRepository{}
}

func (r *stockRepository) GetStock(ticker string) (*models.Stock, error) {
	time.Sleep(time.Duration(rand.Intn(400)+100) * time.Millisecond)

	// random price
	price := 10.0 + rand.Float64()*490.0

	return &models.Stock{
		Ticker:     ticker,
		Price:      price,
		LastUpdate: time.Now(),
	}, nil
}
