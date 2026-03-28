package repositories

import (
	"context"
	"stock-board-go/internal/models"
)

type StockRepository interface {
	GetStock(ctx context.Context, ticker string) (*models.Stock, error)
}

type stockRepository struct {
	stock *models.Stock
}

func NewStockRepository() StockRepository {
	return &stockRepository{}
}

func (r *stockRepository) GetStock(ctx context.Context, ticker string) (*models.Stock, error) {
	return r.stock, nil
}
