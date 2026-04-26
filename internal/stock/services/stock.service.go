package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stock-board-go/internal/stock/models"
	"stock-board-go/internal/stock/repositories"
	"time"
)

type StockService interface {
	GetStock(symbol string) (models.Stock, error)
}

type stockService struct {
	apiUrl    string
	apiKey    string
	client    *http.Client
	stockRepo repositories.StockRepository
}

func NewStockService(apiUrl, apiKey string, stockRepo repositories.StockRepository) StockService {
	return &stockService{
		apiUrl:    apiUrl,
		apiKey:    apiKey,
		client:    &http.Client{Timeout: 10 * time.Second},
		stockRepo: stockRepo,
	}
}

func (s *stockService) GetStock(symbol string) (models.Stock, error) {
	url := fmt.Sprintf("%s?symbol=%s&token=%s", s.apiUrl, symbol, s.apiKey)
	resp, err := s.client.Get(url)
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

	stock := models.Stock{
		Ticker:     symbol,
		Price:      result.CurrentPrice,
		LastUpdate: time.Now(),
	}

	if err := s.stockRepo.SaveStock(&stock); err != nil {
		log.Println("Error Saving:", symbol, err)
		return models.Stock{}, err
	}

	return stock, nil
}
