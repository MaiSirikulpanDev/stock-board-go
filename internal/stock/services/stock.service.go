package services

import (
	"stock-board-go/internal/stock/models"
	"stock-board-go/internal/stock/repositories"
)

type StockService interface {
	GetStock(ticker string) (*models.Stock, error)
}

type stockService struct {
	stockRepo repositories.StockRepository
}

func NewStockService(stockRepo repositories.StockRepository) StockService {
	return &stockService{stockRepo: stockRepo}
}

func (s *stockService) GetStock(ticker string) (*models.Stock, error) {
	return s.stockRepo.GetStock(ticker)
}