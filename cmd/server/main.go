package main

import (
	"net/http"
	"stock-board-go/internal/stock/controllers"
	"stock-board-go/internal/stock/repositories"
	"stock-board-go/internal/stock/services"
)

func main() {
	stockRepo := repositories.NewStockRepository()
	stockService := services.NewStockService(stockRepo)
	stockController := controllers.NewStockController(stockService)

	http.HandleFunc("/ws/stock", stockController.GetStock)
	http.ListenAndServe(":8080", nil)
}
