package main

import (
	"fmt"
	"log"
	"os"
	"stock-board-go/internal/stock/controllers"
	"stock-board-go/internal/stock/models"
	"stock-board-go/internal/stock/repositories"
	"stock-board-go/internal/stock/services"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func fetchAndPrintStocks(symbols []string, stockController controllers.StockController) {
	var wg sync.WaitGroup
	stockChannel := make(chan models.Stock, len(symbols))

	for _, sym := range symbols {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			if stock, err := stockController.GetStock(s); err == nil {
				stockChannel <- stock
			} else {
				log.Printf("Failed to get stock %s: %v", s, err)
			}
		}(sym)
	}

	wg.Wait()
	close(stockChannel)

	for stock := range stockChannel {
		fmt.Printf("Ticker: %s, Price: %.2f, Last Update: %s\n", stock.Ticker, stock.Price, stock.LastUpdate.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("------------------------------------------------")
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	apiUrl := os.Getenv("API_URL")
	apiKey := os.Getenv("FINNHUB_API_KEY")
	symbolStr := os.Getenv("SYMBOL")

	if apiUrl == "" || apiKey == "" || symbolStr == "" {
		log.Fatal("Missing required environment variables: API_URL, FINNHUB_API_KEY, or SYMBOL")
	}

	stockRepo := repositories.NewStockRepository()
	stockService := services.NewStockService(apiUrl, apiKey, stockRepo)
	stockController := controllers.NewStockController(stockService)

	symbols := strings.Split(symbolStr, ",")

	fmt.Println("Live Stock Price:")
	fetchAndPrintStocks(symbols, stockController)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fetchAndPrintStocks(symbols, stockController)
	}
}
