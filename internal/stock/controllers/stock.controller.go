package controllers

import (
	"encoding/json"
	"net/http"
	"stock-board-go/internal/stock/services"
)

type StockController interface {
	GetStock(w http.ResponseWriter, r *http.Request)
}

type stockController struct {
	stockService services.StockService
}

func NewStockController(stockService services.StockService) StockController {
	return &stockController{stockService: stockService}
}

func (c *stockController) GetStock(w http.ResponseWriter, r *http.Request) {
	ticker := r.URL.Query().Get("ticker")
	if ticker == "" {
		http.Error(w, "ticker is required", http.StatusBadRequest)
		return
	}

	stock, err := c.stockService.GetStock(ticker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}