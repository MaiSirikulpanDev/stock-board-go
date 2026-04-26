package controllers

import (
	"stock-board-go/internal/stock/models"
	"stock-board-go/internal/stock/services"
	"sync"
)

type StockController interface {
	GetStock(symbol string, wg *sync.WaitGroup, c chan<- models.Stock)
}

type stockController struct {
	stockService services.StockService
}

func NewStockController(stockService services.StockService) StockController {
	return &stockController{stockService: stockService}
}

func (controller *stockController) GetStock(symbol string, wg *sync.WaitGroup, c chan<- models.Stock) {
	defer wg.Done()

	stock, err := controller.stockService.GetStock(symbol)
	if err != nil {
		return
	}

	c <- stock
}
