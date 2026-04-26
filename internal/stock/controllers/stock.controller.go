package controllers

import (
	"stock-board-go/internal/stock/models"
	"stock-board-go/internal/stock/services"
)

type StockController interface {
	GetStock(symbol string) (models.Stock, error)
}

type stockController struct {
	stockService services.StockService
}

func NewStockController(stockService services.StockService) StockController {
	return &stockController{stockService: stockService}
}

func (controller *stockController) GetStock(symbol string) (models.Stock, error) {
	return controller.stockService.GetStock(symbol)
}

