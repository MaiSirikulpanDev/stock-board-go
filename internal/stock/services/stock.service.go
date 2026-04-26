package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"stock-board-go/internal/stock/models"
	"stock-board-go/internal/stock/repositories"
)

type StockService interface {
	GetStock(symbol string) (models.Stock, error)
}

type stockService struct {
	stockRepo repositories.StockRepository
}

func NewStockService(stockRepo repositories.StockRepository) StockService {
	return &stockService{stockRepo: stockRepo}
}

func (s *stockService) GetStock(symbol string) (models.Stock, error) {
	apiUrl := os.Getenv("API_URL")
	url := fmt.Sprintf("%s?symbol=%s&token=%s", apiUrl, symbol, os.Getenv("FINNHUB_API_KEY"))
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error Fetching:", symbol, err)
		return models.Stock{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Error Fetching:", symbol, resp.Status)
		return models.Stock{}, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	var result struct {
		CurrentPrice float64 `json:"c"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error Decoding:", symbol, err)
		return models.Stock{}, err
	}

	stock, err := s.stockRepo.GetStock(symbol, result.CurrentPrice)
	if err != nil {
		return models.Stock{}, err
	}

	return *stock, nil
}
